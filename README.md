# Golang Laboratory [![License](https://img.shields.io/badge/license-BSD%202--Clause-green.svg)](https://opensource.org/licenses/BSD-2-Clause) [![Build Status](https://github.com/rbee3u/golab/actions/workflows/build.yml/badge.svg)](https://github.com/rbee3u/golab/actions?query=branch%3Amain)

### CPU与缓存基础训练

```
1. 运行 github.com/rbee3u/golab/baseline/simple/main.go，查看运行结果。
   目标：记住这种情况下每秒大概可以执行多少条指令，并且能知道具体的原因。
   提示：CPU架构、CPU和Core(核)的区别、什么是时钟频率、什么是CPU主频。

2. 运行 github.com/rbee3u/golab/baseline/memory/main.go，查看运行结果。
   目标：记住这种情况下每秒大概可以执行多少条指令，并且能知道具体的原因。
   提示：CPU缓存、L1、L2、L3分别的大小、程序运行时访问它们的具体速度。

3. 运行 github.com/rbee3u/golab/baseline/atomic/main.go，查看运行结果。
   目标：记住这种情况下每秒大概可以执行多少条指令，并且能知道具体的原因。
   提示：内存顺序(Memory Order)、CPU缓存一致性、内存屏障、Golang的atomic实现。

4. 运行 github.com/rbee3u/golab/baseline/lock/main.go，查看运行结果。
   目标：记住这种情况下每秒大概可以执行多少条指令，并且能知道具体的原因。
   提示：内存顺序(Memory Order)、CPU缓存一致性、内存屏障、Golang的lock实现。
```

### CPU与缓存进阶训练

```
1. 运行 github.com/rbee3u/golab/parallel_optimization/main.go，查看运行结果。
   目标：fast比slow大致快多少，分析背后的原因。
   提示：并行与并发的区别，哪个才可以真正加速程序。

2. 运行 github.com/rbee3u/golab/instruction_pipeline/main.go，查看运行结果。
   目标：fast比slow大致快多少，分析背后的原因。
   提示：什么是指令流水线，什么情况下具有较好的加速效果。

3. 运行 github.com/rbee3u/golab/branch_prediction/main.go，查看运行结果。
   目标：fast比slow大致快多少，分析背后的原因。
   提示：分支预测，它和指令流水线有什么联系。

4. 运行 github.com/rbee3u/golab/spatial_locality/main.go，查看运行结果。
   目标：fast比slow大致快多少，分析背后的原因。
   提示：CPU缓存，程序的局部性原理。

5. 运行 github.com/rbee3u/golab/false_sharing/main.go，查看运行结果。
   目标：fast比slow大致快多少，分析背后的原因。
   提示：CPU缓存，内存屏障，False Sharing。

6. 运行 github.com/rbee3u/golab/lock_sharding/main.go，查看运行结果。
   目标：fast比slow大致快多少，分析背后的原因。
   提示：CPU缓存，内存屏障，分片为什么能够加速，背后的概率学原理是什么。

7. 运行 github.com/rbee3u/golab/bounds_check_elimination/main.go，查看运行结果。
   目标：fast比slow大致快多少，分析背后的原因。
   提示：Bounds Check Elimination，Golang代码如何生成汇编，如何分析汇编得出结论。
```
