"""
标识符 - 唯一性关系链，将【关系链】转换成【中文表述】
"""

import re
from relationship.data.pair import pair as _pair
from relationship.data.sort import sort as _sort
from .unit import number2zh


# 逆转ID
def reverse_id(id_str: str, sex: int):
    if not id_str:
        return [""]
    # 映射关系
    hash_map = {
        "f": ["d", "s"],
        "m": ["d", "s"],
        "h": ["w", ""],
        "w": ["", "h"],
        "s": ["m", "f"],
        "d": ["m", "f"],
        "lb": ["os", "ob"],
        "ob": ["ls", "lb"],
        "xb": ["xs", "xb"],
        "ls": ["os", "ob"],
        "os": ["ls", "lb"],
        "xs": ["xs", "xb"],
    }
    # 年纪判断
    age = ""
    if id_str.endswith("&o"):
        age = "&l"
    elif id_str.endswith("&l"):
        age = "&o"
    id_str = re.sub(r"&[ol\d+]", "", id_str)
    # 性别判断
    if sex < 0:
        if id_str.startswith("w"):
            sex = 1
        elif id_str.startswith("h"):
            sex = 0

    def doing(sex: int):
        sid = "," + str(sex) + "," + id_str
        sid = re.sub(r",[fhs]|,[olx]b", ",1", sid)
        sid = re.sub(r",[mwd]|,[olx]s", ",0", sid)
        sid = sid[:-2]
        sid_arr = sid.split(",")[::-1]
        r_id = ",".join(
            [
                hash_map[id_part][int(sid_arr[i])]
                for i, id_part in enumerate(id_str.split(",")[::-1])
            ]
        )
        gen = get_gen_by_id(r_id)
        return r_id + ("" if gen else age)

    if sex < 0:
        return [doing(1), doing(0)]
    else:
        return [doing(sex)]


# ID列表去重
def filter_id(arr: list[str]):
    def repl(item):
        repl_item = re.sub(r"[ol](?=[s|b])", "x", item)
        repl_item = re.sub(r"&[ol]", "", repl_item)
        return repl_item

    same_list = [item for item in arr if repl(item) == item]

    def filter_item(item):
        repl_item = repl(item)
        return item in same_list or (item != repl_item and repl_item not in same_list)

    filtered = filter(filter_item, arr)

    return list(dict.fromkeys(filtered))  # 去重保持顺序


# 通过ID获取世代数
def get_gen_by_id(id_str: str):
    g_map = {"f": 1, "m": 1, "s": -1, "d": -1}
    gen = 0
    for sub in id_str.split(","):
        s = re.sub(r"&[ol\d]+", "", sub)
        gen += g_map.get(s, 0)
    return gen


# 通过ID获取关系称呼
def get_items_by_id(id_str: str, mode_data: dict[str, list[str]]):
    items = []

    def get_data(key):
        ids = []
        k1 = re.sub(r"(,[sd])(,[wh])?$", r"\1&o\2", key)
        k2 = re.sub(r"(,[sd])(,[wh])?$", r"\1&l\2", key)
        if k1 in mode_data and k2 in mode_data:
            ids = [k1, k2]
        elif key in mode_data:
            ids = [key]
        return [mode_data[id_val][0] for id_val in filter_id(ids)]

    # 对排序进行处理
    match = re.search(r"&([\d]+)(,[hw])?$", id_str)
    if match:
        num = match.group(1)
        zh = number2zh(num)
        id_str = re.sub(r"&\d+", "", id_str)
        if id_str in _sort:
            item = _sort[id_str][0].replace("几", zh)
            items.append(item)
        elif id_str in mode_data:
            gen = get_gen_by_id(id_str)
            item = ""
            if gen < 3 and not re.search(r"[hw],", id_str):
                for name in mode_data[id_str]:
                    if not item and "几" in name:
                        item = name.replace("几", zh)
                if not item:
                    item = mode_data[id_str][0]
                    if re.match(r"^[大小]", item):
                        item = re.sub(r"^[大小]", zh, item)
                    else:
                        item = zh + item
            items.append(item)

    # 直接匹配称呼
    if not items:
        id_str = re.sub(r"&\d+", "", id_str)
        items = get_data(id_str)

    # 忽略年龄条件查找
    if not items:
        id_str = re.sub(r"&[ol]", "", id_str)
        items = get_data(id_str)

    # 忽略年龄条件查找
    if not items:
        id_str = re.sub(r"[ol](b|s)", r"x\1", id_str)
        items = get_data(id_str)

    # 缩小访问查找
    if not items:
        l_id = id_str.replace("x", "l")
        o_id = id_str.replace("x", "o")
        items = get_data(o_id) + get_data(l_id)

    return items


# 通过ID获取关系链
def get_chain_by_id(id_str: str, mode_data: dict[str, list[str]], sex=-1):
    data = {**mode_data, "xb": ["兄弟"], "xs": ["姐妹"]}
    item = "的".join(
        [data[re.sub(r"&[ol\d]+", "", sign)][0] for sign in id_str.split(",")]
    )

    if sex > -1 and f"{sex},{id_str}" in data:
        if sex == 0:
            item = "(女性)" + item
        elif sex == 1:
            item = "(男性)" + item

    return item


# 通过ID获取关系合称
def get_pairs_by_id(id1: str, id2: str):
    from .selector import selector2id

    id1 = re.sub(r"&\d+", "", id1)
    id2 = re.sub(r"&\d+", "", id2)
    id1_x = re.sub(r"([ol])([bs])", r"x\2", id1)
    id2_x = re.sub(r"([ol])([bs])", r"x\2", id2)
    id1_r = re.sub(r"&[ol]", "", id1)
    id2_r = re.sub(r"&[ol]", "", id2)

    def repl_every_item(selector_list: list[str]):
        return [
            re.sub(r"([ol])([bs])", r"x\2", re.sub(r"&[ol\d]+", "", selector))
            for selector in selector_list
        ]

    result = []
    result_x = []
    result_r = []
    for key in _pair:
        selectors = key.split("#")
        if len(selectors) > 1:
            list1 = selector2id(selectors[0])
            list2 = selector2id(selectors[1])
            list1_r = repl_every_item(list1)
            list2_r = repl_every_item(list2)

            if (id1 in list1 and id2 in list2) or (id2 in list1 and id1 in list2):
                result.append(_pair[key][0])

            if (id1_x in list1_r and id2_x in list2_r) or (
                id2_x in list1_r and id1_x in list2_r
            ):
                result_x.append(_pair[key][0])

            if (id1_r in list1_r and id2_r in list2_r) or (
                id2_r in list1_r and id1_r in list2_r
            ):
                result_r.append(_pair[key][0])

    if not result:
        result = result_x

    if not result:
        result = result_r

    return result
