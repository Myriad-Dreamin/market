# Minimum Market
market backend powered by minimum

## interface definition
https://github.com/Myriad-Dreamin/market/blob/master/control

## Documentation
API文档：

https://myriaddreamin.docs.apiary.io/

https://github.com/Myriad-Dreamin/market/blob/master/apiary.apib

## Installation

如果你需要在物理机上编译，那么你需要先下载[golang](https://golang.org/dl/)和[gcc](https://gcc.gnu.org/)，并设置下列环境变量：

```bash
export GO111MODULE=on GOPROXY=https://goproxy.io
```

原生的安装方法如下：

```bash
go install github.com/Myriad-Dreamin/go-magic-package/package-attach-to-path
```



## Minimum-Cli/Makefile

python环境如下：

```bash
$ python --version
Python 2.7.15
$ python3 --version
Python 3.7.2
```

使用`cli.py`需要安装`fire`

```bash
pip install fire
```

`python-fire`的文档参见[python-fire](https://github.com/google/python-fire)

如果你仅仅安装，则不需要安装`fire`，因为`makefile.py`仅仅依赖于`pymake.py`，而`pymake`只需要原生的`python2/3`模块。

`cli`有如下命令：

```plain
NAME                                      
    cli.py                                
                                          
SYNOPSIS                                  
    cli.py COMMAND | VALUE                
                                          
COMMANDS                                  
    COMMAND is one of the following:      
                                          
     apply_context                        
                                          
     create_pure_service                  
                                          
     create_router                        
                                          
     create_service                       
                                          
     create_template                      
                                          
     fast_generate                        
                                          
     generate                             
                                          
     hello                                
                                          
     install                              
                                          
     make                                 
                                          
     replace                              
                                          
     template_to                          
                                          
     templates_to                         
                                          
     test
```

#### apply_context

```bash
NAME
    cli.py apply_context

SYNOPSIS
    cli.py apply_context <flags> [KEY_VALUES]...

POSITIONAL ARGUMENTS
    KEY_VALUES

FLAGS
    --f=F (required)
    --c=C (required)
```

此命令用于帮助设置`makefile`的context，context保存在`.market-env.json`中，你可以手动修改。

