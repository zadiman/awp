awp is a tool to parse *($HOME/.aws/config)* and let you choose via a fuzzy finder what profile to use and create a way to source aws profiles into your env. When run a file will get created `/tmp/aws_profile` with two export statements:
```
$ cat /tmp/aws_profile 
export AWS_PROFILE=aws-profile
export AWS_REGION=eu-north-1
```

## Install
```
$ git clone git@github.com:zadiman/awp.git
$ cd awp
$ go build
$ sudo cp awp /usr/local/bin
```
Now create an alias in your `.bashrc` or `.zshrc`
```
alias awp="awp && . /tmp/aws_profile"
```

## Usage
Now when you execute your alias:

![image](https://user-images.githubusercontent.com/26366265/169647564-7010f98d-ac70-4440-94b4-94f7b8a73ae7.png)

## Additional use
Choose what region to populate `AWS_REGION` with.

```
$ awp -region eu-north-1
```
