"""
模式数据
"""

from .options import OptionMode
from .map import init_map


# 获取模式数据
def get_mode_data(mode: OptionMode):
    data = {**init_map()}
    if mode == OptionMode.GUANG_DONG:
        from relationship.locale_ import guangdong

        for k, v in guangdong["data"].items():
            data[k] = v + data.get(k, [])

    elif mode == OptionMode.NORTH:
        from  relationship.locale_ import north

        for k, v in north["data"].items():
            data[k] = v + data.get(k, [])

    return data
