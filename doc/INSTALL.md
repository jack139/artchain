## 安装和启动

### 编译
```bash
GOPATH=$HOME/Codes/go starport build
```

### 第一次启动准备
```bash
bash shell/init_chain.sh
```
修改 n1/config/app.toml 中 api 小节中，enable=true

### 启动
```bash
artchaind start --log_level warn --home n1
```

### 启动http服务
```bash
artchaind http 8888 --chain-id artchain --home n1 --from user0
```

### 使用starport启动
```bash
GOPATH=$HOME/Codes/go starport serve --home n1
```

### 安装ipfs
1. 第一个节点
```bash
./install.sh
```
查看ipfs id，修改bootstrap.txt内容里的id和ip

2. 其他节点
复制新的bootstrap.txt，然后执行 install.sh

3. 在各个节点
```bash
./run.sh
```

4. 查看节点状态
```bash
ipfs swarm peers
```

5. 测试节点

节点1
```bash
# echo 'ipfs1' > ipfs1.txt
# ipfs add ipfs1.txt
added QmZyTztEF1UfJ1Qw8HzaWpQcv98ogu4kfnKteUwav2gg6T ipfs1.txt
 6 B / 6 B [==========================] 100.00%
```

节点2
```bash
# ipfs get QmZyTztEF1UfJ1Qw8HzaWpQcv98ogu4kfnKteUwav2gg6T
Saving file(s) to QmZyTztEF1UfJ1Qw8HzaWpQcv98ogu4kfnKteUwav2gg6T
 6 B / 6 B [==========================] 100.00% 0s
```
