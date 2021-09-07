# Git 全局设置:
```shell
git config --global user.name "骆昊宇"
git config --global user.email "202806609@qq.com"
```

# 创建 git 仓库:
```shell
mkdir gin-result
cd gin-result
git init
touch README.md
git add README.md
git commit -m "first commit"
git remote add origin https://gitee.com/luohaoyu/gin-result.git
git push -u origin master
```

# 已有仓库?
```shell
cd existing_git_repo
git remote add origin https://gitee.com/luohaoyu/gin-result.git
git push -u origin master
```