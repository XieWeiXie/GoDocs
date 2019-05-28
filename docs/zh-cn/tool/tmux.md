# Tmux

> 键盘回显工具：KeyCastr

### 1.TMUX

终端复用工具

概念：

- server
- session
- window
- pane


### 2. 快捷键

前缀：`ctrl + b`


```
? : 帮助
ctrl + o: 窗口调换位置
o: 下个窗口
空格: 调整布局
！: 当前窗口为新窗口
“ : 横向分割
% : 纵向分割
q : 显示分割窗口编号
n : 下个窗口
t : 显示时间
c : 创建新窗口
0-9: 窗口
l : 最后的窗口
p : 前一窗口
w : 以菜单的方式显示和选择窗口
x : 关闭窗口
& : 关闭窗口
s : 以菜单的方式显示和选择窗口
d : 退出 tumx
: : 以命令的方式进行 tmux 配置

```

```
tmux new -s session
tmux new -s session -d
tmux ls
tmux attach -t session
tmux kill-session -t basic // 关闭
```

```
:set -g mouse on 滚动屏幕
```