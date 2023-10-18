## 编译电路
```azure
circom mul.circom --r1cs --wasm
```
或者
```azure
circom mul.circom --r1cs --wasm --sym --c
```

## 可信设置
```azure
 //开启可信设置仪式。生成pot12_0000.ptau
snarkjs powersoftau new bn128 12 pot12_0000.ptau -v

//贡献一次随机数（可多次），生成pot12_0001.ptau文件
snarkjs powersoftau contribute pot12_0000.ptau pot12_0001.ptau --name="First contribution" -v

//进入设置仪式第二阶段，生成pot12_final.ptau文件
snarkjs powersoftau prepare phase2 pot12_0001.ptau pot12_final.ptau -v

//利用pot12_final.ptau和r1cs约束文件，生成mul_0000.zkey（类似于私钥）
snarkjs groth16 setup mul.r1cs pot12_final.ptau mul_0000.zkey

//对zkey贡献一次随机数（可多次），生成mul_0001.zkey文件
snarkjs zkey contribute mul_0000.zkey mul_0001.zkey --name="1st Contributor Name" -v

//导出verification_key（类似于公钥）
snarkjs zkey export verificationkey mul_0001.zkey verification_key.json

```

## 生成witness
生成可信文件
```azure
cd mul_js
vim input.json

//input.json中填写下面内容，对应我们的电路代码Mul，输入a=3, b=11
{"a": "3", "b": "11"}
```
利用电路代码生成witness
```azure
node generate_witness.js mul.wasm input.json witness.wtns
```

## 检查验证
```azure
//检查r1cs文件，可以查看constrains数目
 snarkjs r1cs info mul.r1cs
 
 //打印constrain
snarkjs r1cs print mul.r1cs mul.sym

//把r1cs输出为json文件，可以直接查看json文件内容
snarkjs r1cs export json mul.r1cs mul.r1cs.json

// 验证witness
snarkjs wtns check mul.r1cs mul_js/witness.wtns
```
## 生成proof
```azure
snarkjs groth16 prove mul_0001.zkey mul_js/witness.wtns proof.json public.json
```

## 验证proof
这一步利用verification_key验证proof.json和执行结果public.json是否合法。
```azure
snarkjs groth16 verify verification_key.json public.json proof.json
```

## 生成合约
```
snarkjs zkey export solidityverifier mul_0001.zkey verifier.sol
```