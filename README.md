# alGorithms

A workspace to manually learn algorithm implementations.

Doing this because I believe in first-principles and understanding fundamentals, regardless of AI agents. So here the agent write the docs, specs, test suite, and verification, while I try my best to implement each algorithm manually.

It feels pretty cool. I get to understand how seemingly opaque and mysterious arts work, and it improves my foundational understanding considerably. Not to mention how fun this is.

An algorithm is done only when both boxes are checked:
`go test ./<module> -run <Name> -v` (add `-out` for tests that produce artifacts) must pass for correctness and complexity.

## sorting

| Algorithm | Difficulty | Done |
|---|---|---|
| BubbleSort | easy | ✓ |
| InsertionSort | easy | ✓ |
| SelectionSort | easy | ✓ |
| MergeSort | medium | - |
| QuickSort | medium | - |
| HeapSort | medium | ✓ |
| RadixSort | hard | - |
| TimSort | hard | - |

## scanning

| Algorithm | Difficulty | Done |
|---|---|---|
| LinearSearch | easy | ✓ |
| BinarySearch | easy | ✓ |
| Kadane | medium | - |
| TwoPointerDedup | medium | ✓ |
| BoyerMooreMajority | medium | - |
| KMP | hard | - |
| RabinKarp | hard | - |

## hashing

| Algorithm | Difficulty | Done |
|---|---|---|
| FNV1a | easy | ✓ |
| ChainMap | medium | - |
| OpenMap | medium | - |
| BloomFilter | hard | - |
| Ring | hard | - |

## numerics

| Algorithm | Difficulty | Done |
|---|---|---|
| KahanSum | easy | - |
| NeumaierSum | easy | - |
| NearEqual | easy | - |
| SafeLog/SafeExp | easy | - |
| Horner | easy | - |
| Welford | medium | - |
| TaylorExp | medium | - |
| LogSumExp | medium | - |

## statistics

| Algorithm | Difficulty | Done |
|---|---|---|
| DiscreteExpectation | easy | - |
| BayesPosterior | easy | - |
| BinomialPMF | easy | - |
| CumulativeDistribution | medium | - |
| KDE1D | medium | - |
| EMBinomial | hard | - |
| EMMultinomial | hard | - |

## calculus_scalar

| Algorithm | Difficulty | Done |
|---|---|---|
| Bisection | easy | - |
| NewtonRoot | easy | - |
| FixedPoint | easy | - |
| GradientDescent | easy | - |
| CentralDifference | easy | - |
| SecantMethod | medium | - |
| NewtonMinimize | medium | - |
| GoldenSection | medium | - |
| MMQuartic | hard | - |

*Note* : All derivatives are hardcoded in the test suite.

## sampling


| Algorithm | Difficulty | Done |
|---|---|---|
| InverseDiscrete | medium | - |
| SampleExponential | medium | - |
| SampleTriangular | medium | - |
| BoxMuller | medium | - |
| RejectionSample | medium | - |
| MetropolisHastings1D | hard | - |

## images

| Algorithm | Difficulty | Done |
|---|---|---|
| FloodFill | easy | - |
| Rotate90/FlipH | easy | - |
| BoxBlur | easy | - |
| BresenhamLine | medium | - |
| MidpointCircle | medium | - |
| FloydSteinbergDither | hard | - |

## audio

| Algorithm | Difficulty | Done |
|---|---|---|
| SineGen | easy | - |
| Normalize | easy | - |
| MixClipped | easy | - |
| Echo | easy | - |
| LinearResample | medium | - |
| MovingAverageLowpass | medium | - |
