import unittest
import time
from relationship import Relationship
import sys
import gc

class TestRelationshipBenchmark(unittest.TestCase):
    def setUp(self):
        self.relationship = Relationship()

    def test_benchmark_simple_queries(self):
        start_time = time.time()
        for _ in range(100):
            self.relationship.relationship({"text": "爸爸的哥哥的弟弟的儿子"})
        elapsed_time = time.time() - start_time
        self.assertLess(elapsed_time, 5.0, f"简单查询性能测试：100次查询耗时 {elapsed_time:.2f} 秒")
        print(f"简单查询性能测试：100次查询耗时 {elapsed_time:.2f} 秒")

    def test_benchmark_complex_queries(self):
        start_time = time.time()
        for _ in range(50):
            self.relationship.relationship({"text": "姐姐的老公的姐姐的老公"})
            self.relationship.relationship({"text": "爸爸的妈妈的姐姐的儿子"})
            self.relationship.relationship({"text": "老婆的外孙的姥爷", "sex": 1})
        elapsed_time = time.time() - start_time
        self.assertLess(elapsed_time, 5.0, f"复杂查询性能测试：150次查询耗时 {elapsed_time:.2f} 秒")
        print(f"复杂查询性能测试：150次查询耗时 {elapsed_time:.2f} 秒")

    def test_benchmark_reverse_queries(self):
        start_time = time.time()
        for _ in range(50):
            self.relationship.relationship({"text": "爸爸的舅舅", "sex": 0, "reverse": True})
            self.relationship.relationship({"text": "岳母", "target": "女儿", "reverse": True})
        elapsed_time = time.time() - start_time
        self.assertLess(elapsed_time, 3.0, f"反向关系查询性能测试：100次查询耗时 {elapsed_time:.2f} 秒")
        print(f"反向关系查询性能测试：100次查询耗时 {elapsed_time:.2f} 秒")

    def test_benchmark_expression_queries(self):
        start_time = time.time()
        for _ in range(50):
            self.relationship.relationship("外婆和奶奶之间是什么关系？")
            self.relationship.relationship("我应该叫外婆的哥哥什么？")
        elapsed_time = time.time() - start_time
        self.assertLess(elapsed_time, 3.0, f"表达式查询性能测试：100次查询耗时 {elapsed_time:.2f} 秒")
        print(f"表达式查询性能测试：100次查询耗时 {elapsed_time:.2f} 秒")

    def test_memory_usage(self):

        # 强制垃圾回收
        gc.collect()

        # 测量初始内存
        initial_size = sys.getsizeof(self.relationship)

        # 执行一系列操作
        for _ in range(20):
            self.relationship.relationship({"text": "爸爸的哥哥的弟弟的儿子"})
            self.relationship.relationship({"text": "姐姐的老公的姐姐的老公"})

        # 再次强制垃圾回收
        gc.collect()

        # 测量最终内存
        final_size = sys.getsizeof(self.relationship)

        # 内存增长不应超过初始大小的50%
        self.assertLess(final_size, initial_size * 1.5,
                         f"内存使用测试：初始大小 {initial_size}kb，最终大小 {final_size}kb")
        print(f"内存使用测试：初始大小 {initial_size}kb，最终大小 {final_size}kb")


if __name__ == "__main__":
    unittest.main(verbosity=2)
