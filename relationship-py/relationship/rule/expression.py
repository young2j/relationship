"""
自然语言表达 - 将【语言表达】转换成【参数调用】
"""

import re
from typing import Callable
from relationship.options import ParameterOptions

expression: list[dict[str, re.Pattern | Callable[[re.Match], ParameterOptions]]] = [
    {
        "exp": re.compile(r"^(\S+?)(应该|得)?(称呼|叫|喊)(\S+?)(什么|啥)？?$"),
        "opt": lambda match: ParameterOptions(
            **{"text": match.group(4), "target": match.group(1)}
        ),
    },
    {
        "exp": re.compile(r"^(\S+?)(应该|得)?(如何|怎么)(称呼|叫|喊)(\S+?)？?$"),
        "opt": lambda match: ParameterOptions(
            **{"text": match.group(5), "target": match.group(1)}
        ),
    },
    {
        "exp": re.compile(r"^(\S+?)是(\S+?)的(谁|什么|什么亲戚|啥亲戚|什么人)？?$"),
        "opt": lambda match: ParameterOptions(
            **{
                "text": match.group(1),
                "target": match.group(2),
                "type": "default" if "的" in match.group(1) else "chain",
            }
        ),
    },
    {
        "exp": re.compile(r"^(\S+?)(应该|得)?管(\S+?)叫(什么|啥)？?$"),
        "opt": lambda match: ParameterOptions(
            **{"text": match.group(3), "target": match.group(1)}
        ),
    },
    {
        "exp": re.compile(r"^(\S+?)(和|与)(\S+?)(之间)?是(什么|啥)关系？?$"),
        "opt": lambda match: ParameterOptions(
            **{
                "text": match.group(1),
                "target": match.group(3),
                "type": "pair",
            }
        ),
    },
    {
        "exp": re.compile(r"^(\S+?)对于(\S+?)是(什么|啥)关系？?$"),
        "opt": lambda match: ParameterOptions(
            **{
                "text": match.group(1),
                "target": match.group(2),
                "type": "chain",
            }
        ),
    },
    {
        "exp": re.compile(r"^(\S+?)(指的)?是(什么|啥)(意思|关系|亲戚关系|辈分|人)？?$"),
        "opt": lambda match: ParameterOptions(
            **{"text": match.group(1), "type": "chain"}
        ),
    },
    {
        "exp": re.compile(r"^(\S+?)(应该|得)?(称呼|叫|喊)(什么|啥)？?$"),
        "opt": lambda match: ParameterOptions(**{"text": match.group(1)}),
    },
    {
        "exp": re.compile(r"^(\S+?)(应该|得)?(如何|怎么)(称呼|叫|喊)？?$"),
        "opt": lambda match: ParameterOptions(**{"text": match.group(1)}),
    },
    {
        "exp": re.compile(
            r"^(\S+?)是(谁|什么|什么亲戚|啥亲戚|什么人|什么辈分|啥辈分)？?$"
        ),
        "opt": lambda match: ParameterOptions(
            **{
                "text": match.group(1),
                "type": "default" if "的" in match.group(1) else "chain",
            }
        ),
    },
]
