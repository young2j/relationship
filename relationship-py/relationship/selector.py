import re
from relationship.rule import filter as _filter
from relationship.rule.replace import replace as _replace
from relationship.rule.similar import similar as _similar
from .unit import zh2number
from .id_ import reverse_id, filter_id, get_gen_by_id
from dataclasses import dataclass


@dataclass
class SelectorOptions(object):
    from_: str
    to: str
    sex: int
    optimal: bool = False


@dataclass
class SelectorResult(object):
    selector: str
    sex: int


# 获得最简关系表达
def get_optimal(options: SelectorOptions):
    from_ = options.from_
    to = options.to
    sex = options.sex

    from_chain = from_.split(",")
    to_chain = to.split(",")

    def get_sex(chain: str):
        if re.match(r"^([fhs1](&[ol\d]+)?|[olx]b)(&[ol\d]+)?", chain):
            return 1
        return 0

    def group1(patterns: list[re.Pattern[str]] = [], string: str = ""):
        for pat in patterns:
            match_: re.Match[str] = pat.search(string)
            if match_:
                return match_.group(1)

        return ""

    for i in range(min(len(from_chain), len(to_chain))):
        if from_chain[i] == to_chain[i]:
            from_ = ",".join(from_chain[i + 1 :])
            to = ",".join(to_chain[i + 1 :])
            sex = get_sex(from_chain[i])
            continue
        else:
            if (
                get_gen_by_id(from_chain[i]) == get_gen_by_id(to_chain[i])
                and re.match(r"^[xol][bs]|^[sd]", from_chain[i])
                and re.match(r"^[xol][bs]|^[sd]", to_chain[i])
            ):
                form_type = re.sub(
                    r"^[xol]([bs])", r"\1", re.sub(r"&([ol\d]+)", "", from_chain[i])
                )
                to_type = re.sub(
                    r"^[xol]([bs])", r"\1", re.sub(r"&([ol\d]+)", "", to_chain[i])
                )

                if form_type != to_type:
                    break

                patterns = [re.compile(r"&([ol\d]+)"), re.compile(r"([ol])[bs]")]
                from_attr = group1(patterns, from_chain[i])
                to_attr = group1(patterns, to_chain[i])

                if from_attr and to_attr:
                    if from_attr.isdigit() and to_attr.isdigit():
                        if int(from_attr) > int(to_attr):
                            from_chain[i] = re.sub(
                                r"^[xol]s|^d",
                                "ls",
                                re.sub(r"^[xol]b|^s", "lb", from_chain[i]),
                            )
                        elif int(from_attr) < int(to_attr):
                            from_chain[i] = re.sub(
                                r"^[xol]s|^d",
                                "os",
                                re.sub(r"^[xol]b|^s", "ob", from_chain[i]),
                            )
                    elif (from_attr.isdigit() and to_attr == "o") or (
                        from_attr == "l" and to_attr.isdigit()
                    ):
                        from_chain[i] = re.sub(
                            r"^[xol]s|^d",
                            "ls",
                            re.sub(r"^[xol]b|^s", "lb", from_chain[i]),
                        )
                    elif (from_attr.isdigit() and to_attr == "l") or (
                        from_attr == "o" and to_attr.isdigit()
                    ):
                        from_chain[i] = re.sub(
                            r"^[xol]s|^d",
                            "os",
                            re.sub(r"^[xol]b|^s", "ob", from_chain[i]),
                        )

                    from_ = ",".join(from_chain[i:])
                    to = ",".join(to_chain[i + 1 :])
                    sex = get_sex(to_chain[i])

                elif options.optimal:
                    from_attr_match = re.search(r"([xol])[bs]", from_chain[i])
                    from_attr = from_attr_match.group(1) if from_attr_match else ""

                    to_attr_match = re.search(r"([xol])[bs]", to_chain[i])
                    to_attr = to_attr_match.group(1) if to_attr_match else ""

                    if from_attr == "x" or to_attr == "x":
                        from_ = ",".join(from_chain[i + 1 :])
                        to = ",".join(to_chain[i + 1 :])
                        sex = get_sex(from_chain[i])
                        continue
            break

    return SelectorOptions(**{"from_": from_, "to": to, "sex": sex})


# 中文获取选择器
def get_selectors(string: str, cache_data: dict[str, list[str]]):
    # 基础字符串替换
    # 将所有"之"替换为"的"
    string = string.replace("之", "的")
    # 将"吾之xxx"或"吾xxx"模式替换为"xxx"
    string = re.sub(r"吾之?(.+)", r"\1", string)
    # 将"我的xxx"或"我xxx"模式替换为"xxx"
    string = re.sub(r"我的?(.+)", r"\1", string)

    # 惯用口语标准化
    if re.search(r"[^娘婆岳亲]家的?(孩子|儿子|女儿)", string):
        string = re.sub(r"家的?", "的", string)

    string = re.sub(r"(舅|姑)+(爸|父|丈|妈|母)?家的?(哥|姐|弟|妹)+", r"\1表\3", string)
    string = re.sub(r"(舅|姑)+(爸|父|丈|妈|母)?家的?", r"\1表", string)
    string = re.sub(r"(伯|叔)+(父|母)?家的?(哥|姐|弟|妹)+", r"堂\3", string)
    string = re.sub(r"(伯|叔)+(父|母)?家的?", "堂", string)
    string = re.sub(r"姨+(爸|父|丈|妈|母)?家的?(哥|姐|弟|妹)+", r"姨\2", string)
    string = re.sub(r"姨+(爸|父|丈|妈|母)?家的?", "姨", string)

    lists = string.split("的")
    result = []
    is_match = True

    def get_list(name: str, keywords: list[str] = []):
        # name == ""
        if not name:
            return

        # 词义扩展
        for item in _replace:
            for word in item["arr"]:
                name1 = re.sub(item["exp"], word, name)
                if name1 != name:
                    keywords.append(name1)
                    get_list(name1, keywords)

        # 同义词替换
        for word, similar in _similar.items():
            name1 = name.replace(word, similar)
            if name1 != name:
                keywords.append(name1)

            name2 = name.replace(similar, word)
            if name2 != name:
                keywords.append(name2)

        return

    while lists:
        name = lists.pop(0)  # 当前匹配词
        items = []  # 当前匹配词可能性
        keywords = [name]
        get_list(name, keywords)

        # 通过关键词找关系
        items_map = [[], [], []]
        for kw in keywords:
            kw = re.sub(r"^[尕幺细满碎晚末尾幼]", "小", kw)
            match = re.match(r"^[大小]|^[一二三四五六七八九十]+", kw)

            if match:  # 匹配排序
                x_name = re.sub(match.group(0), "几", kw)
                r_name = re.sub(match.group(0), "", kw)
                num = zh2number(match.group(0))

                for index, item in enumerate([x_name, r_name, kw]):
                    ids = cache_data.get(item, [])
                    for i in ids:
                        id = re.sub(r"(,[hw])$", f"&{num}\\1", i)
                        id = re.sub(r"([^hw]+)$", f"\\1&{num}", id)
                        if not re.match(r"^[mf,]+$", i) and not re.match(
                            r"^[从世]", kw
                        ):  # 直系祖辈不参与排序
                            items_map[index].append(id)

            items.extend(cache_data.get(kw, []))

        # 如找不到结果，检查是否存在称呼的排行问题
        for items_x in items_map:
            if not items:
                items = items_x

        # 完全匹配不到结果
        if not items:
            is_match = False

        res = []
        if not result:
            result = [""]

        for a in result:
            for b in items:
                res.append(a + ("," + b if b else ""))
        result = res

    return filter_id(result) if is_match else []


# 合并选择器，查找两个对象之间的关系
def merge_selector(param: SelectorOptions) -> list[SelectorResult]:
    _sex = param.sex

    if _sex < 0:
        to_sex = -1
        from_sex = -1
        if re.match(r"^,[w1]", param.from_):
            from_sex = 1
        elif re.match(r"^,[h0]", param.from_):
            from_sex = 0
        if re.match(r"^,[w1]", param.to):
            to_sex = 1
        elif re.match(r"^,[h0]", param.to):
            to_sex = 0

        if from_sex == -1 and to_sex > -1:
            _sex = to_sex
        elif from_sex > -1 and to_sex == -1:
            _sex = from_sex
        elif from_sex == to_sex:
            _sex = from_sex
        else:
            return []

    from_ids = selector2id(param.from_, _sex)
    to_ids = selector2id(param.to, _sex)

    if not from_ids or not to_ids:
        return []

    results: list[SelectorResult] = []
    for from_id in from_ids:
        for to_id in to_ids:
            sex = _sex
            selector = "," + to_id

            if re.search(r",([fhs1](&[ol\d]+)?|[olx]b)(&[ol\d]+)?$", selector):
                sex = 1
            if re.search(r",([mwd0](&[ol\d]+)?|[olx]s)(&[ol\d]+)?$", selector):
                sex = 0

            if from_id and to_id:
                is_optimal = param.optimal
                if re.search(r"&\d+", from_id) or re.search(r"&\d+", to_id):
                    is_optimal = True

                if is_optimal:
                    optimal_options = get_optimal(
                        SelectorOptions(
                            **{
                                "from_": from_id,
                                "to": to_id,
                                "sex": _sex,
                                "optimal": param.optimal,
                            }
                        )
                    )
                    from_id = optimal_options.from_
                    to_id = optimal_options.to
                    _sex = optimal_options.sex

            to_rids = reverse_id(to_id, _sex) if to_id else [""]

            for to_r in to_rids:
                selector = (f",{to_r}" if to_r else "") + (
                    f",{from_id}" if from_id else ""
                )
                results.append(SelectorResult(**{"selector": selector, "sex": sex}))

    return results


# 扩展选择器
def expand_selector(selector: str) -> list[str]:
    """
    扩展选择器，将【关系链】缩减成最简表达形式

    Args:
        selector: 关系选择器字符串

    Returns:
        list[str]: 扩展后的选择器列表
    """
    result = []
    hash_map = {}

    def get_selector(selector: str) -> bool:
        if not hash_map.get(selector):
            hash_map[selector] = True
            s = ""

            while True:
                s = selector
                for item in _filter:
                    # print('[filter]', item['exp'], selector)

                    selector = re.sub(item["exp"], item["str"], selector)
                    if "#" in selector:
                        for part in selector.split("#"):
                            get_selector(part)
                        return False

                if s == selector:
                    break

            # 同志关系去除
            if re.search(r",[mwd0](&[ol\d+])?,w|,[hfs1](&[ol\d]+)?,h", selector):
                return False

            result.append(selector)

    get_selector(selector)

    return result


# 选择器转ID
def selector2id(selector: str, sex: int | None = None) -> list[str]:
    if not selector.startswith(","):
        selector = "," + selector

    # 性别判断
    if sex is not None:
        if sex < 0:
            if re.match(r"^,[w1]", selector):
                sex = 1
            elif re.match(r"^,[h0]", selector):
                sex = 0
        elif sex == 1 and re.match(r"^,[h0]", selector):
            return []
        elif sex == 0 and re.match(r"^,[w1]", selector):
            return []

        if sex > -1 and ",1" not in selector and ",0" not in selector:
            selector = "," + str(sex) + selector

    # 同志关系去除
    if re.search(r",[mwd0](&[ol\d]+)?,w|,[hfs1](&[ol\d]+)?,h", selector):
        return []


    expanded = expand_selector(selector)
    result = [
        re.sub(r",[01]", "", s)[1:] for s in expanded
    ]  # 去前面逗号和性别信息

    return filter_id(result)
