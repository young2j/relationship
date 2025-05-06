"""
缓存数据
"""

from relationship.data.input import input_ as _input
from relationship.data.sort import sort as _sort


def get_cache_data(mode_data: dict[str, list[str]]):
    _hash = {**mode_data}
    for key in _input:
        _hash[key] = _hash.get(key, []) + _input[key]
    for key in _sort:
        _hash[key] = _hash.get(key, []) + _sort[key]

    cache_data: dict[str, list[str]] = {}
    for key in _hash:
        for name in _hash[key]:
            if name not in cache_data:
                cache_data[name] = []
            cache_data[name].append(key)

    return cache_data
