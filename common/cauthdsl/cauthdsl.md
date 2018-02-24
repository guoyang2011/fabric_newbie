1. 实现策略（policy）相关数据结构和方法
2. 策略主要分为签名策略（SignaturePolicy）和隐式策略(ImplicitMetaPolicy)
3. 最常用的签名策略为SignaturePolicyEnvelope结构，定义在protos/common/policies.go
签名策略规则为检测签名是否满足指定的principal(特定个体、身份等)
更通用的Implicit策略是一个递归结构，可以指定依赖其他策略，最终底层为签名策略
