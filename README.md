# Gitee
## Git 全局设置:
```shell
git config --global user.name "骆昊宇"
git config --global user.email "202806609@qq.com"
```

## 创建 git 仓库:
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

## 已有仓库?
```shell
cd existing_git_repo
git remote add origin https://gitee.com/luohaoyu/gin-result.git
git push -u origin master
```

# github
## …or create a new repository on the command line
```shell
echo "# gin-result" >> README.md
git init
git add README.md
git commit -m "first commit"
git branch -M main
git remote add origin https://github.com/3-838/gin-result.git
git push -u origin main
```

## …or push an existing repository from the command line
```shell
git remote add origin https://github.com/3-838/gin-result.git
git branch -M main
git push -u origin main
```

## …or import code from another repository
> You can initialize this repository with code from a Subversion, Mercurial, or TFS project.

# 更换Git地址
```shell
# 查看
git remote -v

# 移除
git remote rm origin

#加入
git remote add origin [地址]

# 最后查看是否更正
git remote -v
```