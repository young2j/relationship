import re
from .mode import get_mode_data
from .options import OptionMode, OptionType, ParameterOptions
from .cache import get_cache_data
from .selector import (
    get_selectors,
    merge_selector,
    selector2id,
    SelectorOptions,
)
from .id_ import (
    get_chain_by_id,
    reverse_id,
    get_pairs_by_id,
    get_items_by_id,
)


class Relationship(object):
    def __init__(self, mode: OptionMode = OptionMode.DEFAULT):
        self.data = {}
        self.cache_data = {}
        self.mode = mode
        self.set_mode(mode)

    def relationship(self, parameter: str | dict | ParameterOptions):
        opts = ParameterOptions.default()
        if isinstance(parameter, ParameterOptions):
            if parameter.mode == OptionMode.NONE:
                parameter.mode = self.mode
            opts = parameter

        elif isinstance(parameter, str):
            opts = ParameterOptions.from_string(
                parameter, mode=self.mode or OptionMode.DEFAULT
            )

        elif isinstance(parameter, dict):
            if not parameter.get("mode"):
                parameter["mode"] = self.mode
            opts = ParameterOptions.from_mapping(parameter)

        if opts.mode != self.mode:
            self.set_mode(opts.mode)

        from_selectors = get_selectors(opts.text, self.cache_data)
        to_selectors = get_selectors(opts.target, self.cache_data)
        if not to_selectors:
            to_selectors = [""]
        # 匹配结果
        result = []
        for from_selector in from_selectors:
            for to_selector in to_selectors:
                merge_datas = merge_selector(
                    SelectorOptions(
                        **{
                            "from_": from_selector,
                            "to": to_selector,
                            "sex": opts.sex,
                            "optimal": opts.optimal,
                        }
                    )
                )
                ids_duplicated = {}
                for merge_data in merge_datas:
                    ids = selector2id(merge_data.selector, merge_data.sex)
                    for id_ in ids:
                        if id_ in ids_duplicated:
                            continue
                        ids_duplicated[id_] = True

                        temps = [id_]
                        sex = merge_data.sex
                        if opts.reverse:
                            temps = reverse_id(id_, sex)
                            if re.search(r"([fhs1](&[ol\d]+)?|[olx]b)$", id_):
                                sex = 1
                            else:
                                sex = 0

                        if opts.type == OptionType.CHAIN:
                            for id_str in temps:
                                item = get_chain_by_id(
                                    id_str, self.data, merge_data.sex
                                )
                                if item:
                                    result.append(item)
                        elif opts.type == OptionType.PAIR:
                            temps = reverse_id(id_, merge_data.sex)
                            for r_id in temps:
                                pairs = get_pairs_by_id(id_, r_id)
                                result.extend(pairs)
                        else:
                            for id_str in temps:
                                items = get_items_by_id(id_str, self.data)
                                if not items:
                                    items = get_items_by_id(
                                        f"{sex},{id_str}", self.data
                                    )
                                result.extend(items)

        return list(dict.fromkeys(result))

    def data_count(self):
        return len(self.data)

    def set_mode(self, mode: OptionMode, extra_data: dict[str, list[str]] = {}):
        self.data = get_mode_data(mode)
        if extra_data:
            for k, v in extra_data.items():
                self.data[k] = v + self.data.get(k, [])

        self.cache_data = get_cache_data(self.data)

    def get_mode_data(self, mode: OptionMode = OptionMode.DEFAULT):
        if mode == self.mode:
            return self.data

        return get_mode_data(mode)


if __name__ == "__main__":
    r = Relationship()
    print(r.relationship("我应该叫外婆的哥哥什么？"))
