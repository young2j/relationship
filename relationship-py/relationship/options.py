"""
通过表达式获取配置
"""

from enum import StrEnum
from dataclasses import dataclass


class OptionMode(StrEnum):
    NONE = ""
    DEFAULT = "default"
    GUANG_DONG = "guang_dong"
    NORTH = "north"


class OptionType(StrEnum):
    DEFAULT = "default"
    CHAIN = "chain"
    PAIR = "pair"


@dataclass
class ParameterOptions(object):
    # 目标对象：目标对象的称谓汉字表达，称谓间用‘的’字分隔
    text: str = ""
    # 相对对象：相对对象的称谓汉字表达，称谓间用‘的’字分隔，空表示自己
    target: str = ""
    # 本人性别：0表示女性,1表示男性
    sex: int = -1
    # 转换类型：'default'计算称谓,'chain'计算关系链,'pair'计算关系合称
    type: OptionType = OptionType.DEFAULT
    # 称呼方式：true对方称呼我,false我称呼对方
    reverse: bool = False
    # 模式选择：使用setMode方法定制不同地区模式，在此选择自定义模式
    mode: OptionMode = OptionMode.DEFAULT
    # 最短关系：计算两者之间的最短关系
    optimal: bool = False

    @classmethod
    def default(cls, **kwargs):
        return cls(**kwargs)

    @classmethod
    def from_mapping(cls, mapping: dict):
        return cls(**mapping)

    @classmethod
    def from_string(cls, text: str, **kwargs):
        from relationship.rule import expression as _expression

        for item in _expression:
            match = item["exp"].match(text)
            if match:
                return item["opt"](match)
        return cls.default(text=text, **kwargs)
