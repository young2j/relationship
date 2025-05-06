"""
数值转换
"""

text_attr: list[str] = ["", "一", "二", "三", "四", "五", "六", "七", "八", "九", "十"]


# 中文数字转阿拉伯数字
def zh2number(text: str):
    num = 0
    map = {"大": 1, "小": 99}
    if text in map:
        num = map[text]
    else:
        text_splits = (
            text.replace("十", "一十", 1) if text.startswith("十") else text
        ).split("十")
        if len(text_splits) == 2:
            dec, unit = text_splits
        else:
            dec, unit = "", text_splits[0]

        unit = text_attr.index(unit)
        dec = text_attr.index(dec)
        num = dec * 10 + unit

    return num


# 阿拉伯数字转中文数字
def number2zh(num: int | str):
    text = ""
    map = {1: "大", 99: "小"}
    if isinstance(num, str):
        num = int(num)

    if num in map:
        text = map[num]
    else:
        dec = num // 10
        unit = num % 10
        text = (text_attr[dec] + "十").replace("一十", "十") if dec else ""
        text += text_attr[unit]

    return text
