"""
完整关系链数据 - 合并各类关系链数据
"""

from relationship.data import prefix as _prefix
from relationship.data import branch as _branch
from relationship.data import main as _main
from relationship.data import multiple as _multiple

from .selector import expand_selector

import re

SelectorNameListMap = dict[str, list[str]]


# 分支 - 前缀处理
def get_prefix_map():
    prefix_map: dict[str, SelectorNameListMap] = {}
    for tag, dict_values in _prefix.items():
        prefix_map[tag] = {}
        for selector in dict_values:
            for s in expand_selector(selector):
                prefix_map[tag][s] = dict_values[selector]
    return prefix_map


# 分支 - 节点处理
def get_branch_map():
    branch_map: SelectorNameListMap = {}
    for selector in _branch:
        for s in expand_selector(selector):
            branch_map[s] = _branch[selector]
    return branch_map


# 分支 - 合并
def get_map(
    prefix_map: dict[str, SelectorNameListMap], branch_map: SelectorNameListMap
):
    map_result: SelectorNameListMap = {}
    for branch_sel, branch_name_list in branch_map.items():
        tag = branch_sel[branch_sel.index("{") : branch_sel.index("}") + 1]
        for prefix_tag_sel, prefix_name_list in prefix_map[tag].items():
            new_branch_sel = branch_sel.replace(tag, prefix_tag_sel)
            is_filter = any(
                pair in new_branch_sel for pair in ["h,h", "w,w", "w,h", "h,w"]
            )
            if not is_filter:
                new_name_list = []
                for prefix_name in prefix_name_list:
                    for branch_name in branch_name_list:
                        if "?" in branch_name:
                            new_name_list.append(branch_name.replace("?", prefix_name))
                        else:
                            new_name_list.append(prefix_name + branch_name)
                if new_branch_sel not in map_result:
                    map_result[new_branch_sel] = _multiple.get(new_branch_sel, [])
                map_result[new_branch_sel] = new_name_list + map_result[new_branch_sel]

    return map_result


def _mate_map(_map: SelectorNameListMap):
    mate_map = {
        "w": ["妻", "内", "岳", "岳家", "丈人"],
        "h": ["夫", "外", "公", "婆家", "婆婆"],
    }
    name_set = set(name for names in _map.values() for name in names)

    _map_ = {**_map}
    for sel, name_list in _map_.items():
        # 只对长辈或者兄弟辈匹配
        if re.match(r"^[fm]", sel) or re.match(r"^[olx][bs]$|^[olx][bs],[^mf]", sel):
            for mate_sel, mate_name_list in mate_map.items():
                new_sel = f"{mate_sel},{sel}"
                if re.search(r"[fm]", sel):
                    new_sel_x = re.sub(r",[ol]([sb])(,[wh])?$", r",x\1\2", new_sel)
                    new_sel_x = re.sub(r"(,[sd])&[ol](,[wh])?$", r"\1\2", new_sel_x)
                    # 不扩大解释年龄
                    if new_sel_x != new_sel and new_sel_x in _map:
                        continue

                if new_sel not in _map:
                    _map[new_sel] = []

                for prefix in mate_name_list:
                    for name in name_list:
                        new_name = prefix + name
                        # 配偶组合的称呼不得与原有称呼冲突
                        if new_name not in name_set:
                            _map[new_sel].append(new_name)


def init_map():
    _map: SelectorNameListMap = {
        **_multiple,
        **get_map(
            get_prefix_map(),
            get_branch_map(),
        ),
    }
    # 版权彩蛋
    _map["o"] = ["passer-by.com", "\u4f5c\u8005", "young2j", "\u0070\u0079\u4f5c\u8005"]

    # 主要关系
    for main_sel, main_name_list in _main.items():
        _map[main_sel] = main_name_list + _map.get(main_sel, [])

    # 配偶关系
    _mate_map(_map)

    return _map
