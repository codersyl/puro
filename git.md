# 不懂就查，可参考的资料
[官方英语文档](https://git-scm.com/docs)
[菜鸟教程](https://www.runoob.com/git/git-basic-operations.html)

# 基本操作

## git init

在当前目录初始化一个仓库

## git add

`git add <file>`把<file>提交到暂存区

## git commit

`git commit -m "<description>"`把所有的暂存区提交到版本库，并添加注释

## git status

`git status`查看当前仓库的状态

## git log

`git log`查看从近到远的日志  
`git log --pretty=oneline`日志显示的更整洁
`git log --graph --pretty=oneline --abbrev-commit`用简洁的图表示

`git reflog`查看之前所有的提交命令以及版本号

## git branch
`git branch` 列出所有分支  

`git branch <branchName>` 创建新的分支branchName

`git branch -d <branchName>` 删除分支branchName

`git branch -D <branchName>` 强制删除分支branchName  
= `git branch --delete --force <branchName>`  

`git branch -m <oldName> <newName>` 把分支名由 oldName 改为 newName   
= `git branch -m <oldName> <newName>`

## git checkout
`git checkout <branchName>` 切换到分支branchName

`git checkout -b <branchName>` 创建分支branchName并切换到分支branchName

## git switch
`git switch <branchName>` 切换到分支branchName

`git switch -c <branchName>` 创建分支branchName并切换到分支branchName

## git reset
`git reset --hard HEAD^^`版本回退2个  

`git reset --hard HEAD~100`版本回退100个  

`git reset --hard <commit_id>`版本回退至<commit_id>版本

`git reset HEAD <file>`可以撤销暂存区某文件的修改

# git remote
`git remote add origin git@github.com:wuhooo/learn-git.git`添加新连接，把本地仓库与github上的仓库连接起来，把远程库叫做origin

`git push -u origin master`把本地仓库的东西推到origin库中，并关联两个库  

`git push origin master`把本地库的 master 分支推送到远程库

`git clone git@github.com:wuhooo/learn-git.git`把远程库弄过来


`git rm <file>`在暂存区中删除<file>，之后还需commit才算真删除

`git checkout -- <file>`让<file>在工作区回退，如果没add，则撤销至与当前版本库一致，如果add过了，则回退至与暂存区中一致

## git merge

`git merge <branch>` 把<branch>合并到现在版本来  
--no-ff参数可以让dev分支不指向与master同一个版本





  

`git stash`把当前工作区的修改隐藏起来  
`git stash list`查看有哪些隐藏的stash  
`git stash apply <stash@{0}>`恢复工作现场  
`git stash drop <stash@{0}>`删除工作现场  
`git stash pop <stash@{0}>`恢复并删除工作现场

`git cherry-pick <commit_id>`重做某一commit的工作，比如在master上修了bug，而当前分支是在bug修复前就分了的，就可以这样操作来在当前分支上修bug  

`git rebase`整理log使得log变成一条线，优点：方便比对；缺点：把分叉改没了  

`git tag`查看所有标签  
`git tag <name>`给最新的commit打标签  
`git tag <name> <commit_id>`给某次提交的版本打标签  
`git show <tagname>`查看标签信息  
`git tag -d <tagname>`删除某标签  
`git push origin <tagname>`把某个标签推到远程  
`git push origin --tags`把所有标签推到远程  
`git push origin :refs/tags/<tagname>`先（另外）删除本地标签，再用这条命令删除远程标签  

`git config --global color.ui true`让git显示颜色  

`git check-ignore`查.gitignore文件中的问题  

`git config --global alias.st status`让`git st`就能实现`git status`的功能  

```
git config --global alias.lg "log --color --graph --pretty=format:'%Cred%h%Creset -%C(yellow)%d%Creset %s %Cgreen(%cr) %C(bold blue)<%an>%Creset' --abbrev-commit"
```

花里胡哨的log配置  

```
$ cat .gitconfig
[alias]
    co = checkout
    ci = commit
    br = branch
    st = status
[user]
    name = Your Name
    email = your@email.com
```

改过的配置都在.gitconifg中，可以用cat查看  

`git remote remove <name>` 删除远程服务器连接


# 常见使用场景

## git新建一个库，连接远程库，拉到本地
```git
git init
git remote add remoteName sshAddress
git fetch remoteName
git merge remoteName/branchNameOfRemote
```
`git fetch remoteName` 相当于把远程库里的东西拉到本地  

`git merge remoteRepName/remoteBranchName` 相当于把本地的origin里的东西更新到本地的git  

可直接`git pull origin remoteBranchName:localBranchName`，相当于 git fetch 加上 git merge 

## git merge冲突
目前出现的情况
* 修改了同一个文件

使用命令 `git merge --no-ff origin/branchName` 来处理冲突，解决后再提交

## 单人开发一次提交流程
```shell
git add 
git commit -m "<description>"
git push <remote_name> <branch_name>
```

## git常规工作流
新开一个分支 feature，开发，提交，push feature。

如果此时main有更新update，  
fetch origin master/main，git rebase main/master，  
进行代码取舍审查，完毕后feature分支就是 初始 -> update -> myChange  
再push到origin，此时相当于对项目管理这发起了一个pull request

项目管理者需要squash and merge

此时本地可以把feature分支删除，然后拉取远程的代码



## 双设备工作流

定期merge，主分支为master。

merge后，设备一拉取最新的master分支，并基于此建立新分支device1，在此分支上进行开发。设备二拉取最新的master分支，并基于此建立新分支device2，在此分支分支上进行开发。

到了下次合并时，设备一把device1分支推送到 origin，设备二把device2分支推送到 origin。选取一个设备把三个分支都从 origin 拉到本地，然后进行如下操作

```git
git switch device1
git merge --no-ff device2
git switch master
git merge --no-ff device1
git branch -d device1
git branch -d device2
git push origin --delete device1
git push origin --delete device2
git push github master
```

此时合并完成，可以在两个设备本地把master之外的分支都删掉，拉取最新的master代码，然后开始下一周期的开发。

## 压缩多次commit

使用 git squash，git squash 不是命令，但它可以把几次提交给压缩起来，下面演示操作流程

假设需要压缩最近5次提交

`git rebase -i HEAD~5` 进入交互式 rebase 模式，Git 会打开一个编辑器，从上到下是从老到新的 5 次提交，形如

```vim
pick <commit1-hash> commit message 1
pick <commit2-hash> commit message 2
pick <commit3-hash> commit message 3
pick <commit4-hash> commit message 4
pick <commit5-hash> commit message 5
```

将下面四个 pick 改为 squash：
```vim
pick <commit1-hash> commit message 1
squash <commit2-hash> commit message 2
squash <commit3-hash> commit message 3
squash <commit4-hash> commit message 4
squash <commit5-hash> commit message 5
````

并按 esc，键入 :wq 回车保存，之后 Git 会再打开一个编辑器，上面是这 5 次 commit 的信息，可以删掉、修改成想要的样子，这些信息会作为一个这几个commit中最早的那次提交的 message 显示在 git log 中，而之后的 4 次提交都会消失，按 esc，键入 :wq 回车保存。

git squash 至此已经完成。

可以使用 `git push -f origin branchName` 推送到远程库。
