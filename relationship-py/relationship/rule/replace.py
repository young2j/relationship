"""
词义扩展 - 对称呼词义进行扩展
"""
replace = [
    {   # 规则1
        'exp': r'^从表',
        'arr': ['从父姑表', '从父舅表', '从父姨表', '从母姑表', '从母舅表', '从母叔表']
    },
    {   # 规则2
        'exp': r'^表表',
        'arr': ['姑表叔表', '姑表姑表', '姑表舅表', '姑表姨表', '舅表叔表', '舅表姑表', '舅表舅表', '舅表姨表']
    },
    {   # 规则3
        'exp': r'^([夫妻内外]?)表',
        'arr': [r'\1姑表', r'\1舅表']
    },
    {   # 规则4
        'exp': r'^([姑舅])表(?=[^伯叔])',
        'arr': [r'\1表伯', r'\1表叔']
    },
    {   # 规则5
        'exp': r'^姻',
        'arr': ['姑姻', '姨姻', '姊妹姻', '女姻']
    },
    {   # 规则6
        'exp': r'^眷',
        'arr': ['叔眷', '舅眷', '兄弟眷', '男眷']
    },
    {   # 规则7
        'exp': r'^亲家',
        'arr': ['姊妹姻', '兄弟眷']
    },
    {   # 规则8
        'exp': r'^([堂表姨]?)([曾高天烈太远鼻]?)(祖?)([伯叔姑舅姨])',
        'arr': [r'\1\4\2\3']
    },
    {   # 规则9
        'exp': r'^([曾高天烈太远鼻]?)祖?王姑',
        'arr': [r'姑\1祖母']
    },
    {   # 规则10
        'exp': r'^([曾玄元来晜仍云耳])([侄甥])',
        'arr': [r'\2\1']
    },
    {   # 规则11
        'exp': r'^外表([伯叔姑舅姨])',
        'arr': [r'姑表\1外', r'舅表\1外']
    },
    {   # 规则12
        'exp': r'([堂表姨]?)外甥',
        'arr': [r'\1甥']
    },
    {   # 规则13
        'exp': r'^([舅叔])([曾玄外]*)孙',
        'arr': [r'\1侄\2孙']
    },
    {   # 规则14
        'exp': r'^([姨姑])([曾玄外]*)孙',
        'arr': [r'\1甥\2孙']
    },
    {   # 规则15
        'exp': r'([孙甥侄])$',
        'arr': [r'\1男', r'\1女']
    },
    {   # 规则16
        'exp': r'([姑舅姨叔])([孙外]*)([男女])$',
        'arr': [r'\1表侄\2\3', r'\1表甥\2\3']
    },
    {   # 规则17
        'exp': r'(.+)父母$',
        'arr': [r'\1父', r'\1母']
    },
    {   # 规则18
        'exp': r'(.+)公婆$',
        'arr': [r'\1公', r'\1婆']
    },
    {   # 规则19
        'exp': r'祖$',
        'arr': ['祖父']
    },
    {   # 规则20
        'exp': r'(新妇|媳)$',
        'arr': ['媳妇']
    },
    {   # 规则21
        'exp': r'嫂$',
        'arr': ['兄妇']
    },
    {   # 规则22
        'exp': r'女儿$',
        'arr': ['女']
    },
]
