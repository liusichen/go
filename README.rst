cli 使用说明
===========

.. Contents :: 目录

安装
----

::
    
    $go build

使用说明
--------

**-h, --help 获取help**

::

    $./cli -h
    $./cli --help

    fast command tools

    fast is a manage system about config file,rpm and etc

    Options:

      -h, --help
          display help information

      -v, --version
          display version

      -l, --list[=false]
          list all sub commands or not

    Commands:

      base-job        create base job discribe
      job             Do job operation such as job creating, status query
      job-template    Do jobtemplate operation
      task            Do task operation such as task creating, status query
      task-template   Do tasktemplate operation


**-v, --version 获取版本**

::

    $./cli -v
    $./cli -version
    v0.0.1

**-l,--list 来查看子命令及其说明**

::

    $./cli -l
    $./cli --list
     base-job       =>  create base job discribe
    job            =>  Do job operation such as job creating, status query
    job-template   =>  Do jobtemplate operation
    task           =>  Do task operation such as task creating, status query
    task-template  =>  Do tasktemplate operation

**注: 如果输入有错会报错**

子命令base-job使用说明
~~~~~~~~~~~~~~~~~~~~

**-h, --help 展示base-job子命令帮助信息(如果./cli base-job也展示帮助信息)**

::

    $./cli base-job -h
    $./cli base-job --help
    $./cli base-job
    create base job discribe

    Options:

      -h, --help
          display help information

      --list
          list all base jobs

      --show
          show the one base job by base job id --show <id>

      --create
          create the base job descriptions --create <jsonfile>

**--list 展示所有base job 描述**

:: 

    $./cli base-job --list
    {
    "code": 0,
    "data": [
      {
        "fields": null,
        "id": 1,
        "name": "Fast-Agent 升级 JOB",
        "type_name": "agent_upgrade"
      },
      {
        "fields": null,
        "id": 2,
        "name": "文件下载 JOB",
        "type_name": "file_download"
      },
      {
        "fields": null,
        "id": 3,
        "name": "RPM 包管理 JOB",
        "type_name": "rpm_mgmt"
      },
      {
        "fields": null,
        "id": 4,
        "name": "SHELL 命令执行 JOB",
        "type_name": " shell_exe"
      },
      {
        "fields": null,
        "id": 5,
        "name": "创建软连接 JOB",
        "type_name": "symlink"
      },
      {
        "fields": null,
        "id": 6,
        "name": "常量文件更新 JOB",
        "type_name": "update_constant"
      }
    ]
     }

**--show <base job id> 展示该id的basejob**

::

    $./cli base-job --show 2
    {
    "code": 0,
    "data": {
      "fields": [
        {
          "cn": "下载地址",
          "en": "url",
          "helper": "文件下载 URL, 支持模板渲染",
          "required": true
        },
        {
          "cn": "MD5 值",
          "en": "md5sum",
          "helper": "文件 MD5 校验字符串, 用于下载后校验",
          "required": true
        },
        {
          "cn": "下载目录",
          "en": "target_dir",
          "helper": "文件下载后的目标目录, 当文件为压缩文件时, 会自动将文件解压放置在此目录",
          "required": true
        },
        {
          "cn": "JOB 超时时间",
          "en": "timeout",
          "helper": "JOB 整体运行超时时间, 包含 pre/post scripts 执行时间",
          "required": true
        },
        {
          "cn": "Pre-Scripts",
          "en": "pre_scripts",
          "helper": "Job 执行前运行的 SHELL 命令, 通常用于环境初始化",
          "required": true
        },
        {
          "cn": "Post-Scripts",
          "en": "post_scripts",
          "helper": "Job 执行后运行的 SHELL 命令, 通常用于执行结果检查",
          "required": true
        }
      ],
      "id": 2,
      "name": "文件下载 JOB",
      "type_name": "file_download"
    }
    }

如果show后面没有参数

:: 

    $./cli base-job --show
    you should input --show <base job id>

job子命令说明
~~~~~~~~~~~~

**-h,--help提示帮助信息**

::

    $./cli job -h
    $./cli job --help
    $./cli job 
    Do job operation such as job creating, status query

    Options:

      -h, --help
          display help information

      --status
          list all job working status

      --filter-by-sn
          list status by SN

      --filter-by-idc
          list status by IDC

      --filter-by-job-id
          list status by JobID

**--status展示所有job状态**

::

    $./cli job --status
    ...展示太多，忽略

    ...

**--filter-by-sn展示某个或者多个sn的状态**

::

    $./cli job --filter-by-sn PR2710R2014122206
    {
        "code": 0,
        "data": [
          {
            "data": "{\"md5sum\":\"6582351ca8d6ac866713afeed5cc9827\",\"post_scripts\":\"-\",\"pre_scripts\":\"-\",\"timeout\":\"10\",\"url\":\"http://10.221.32.13/download/fast/download/fast/fast-agent/fast-agent_1.0.tar.gz\",\"version\":\"1.0\"}",
            "job_id": "6e3fa09ba02483eb40bcf949bf44c157",
            "job_template_id": 5,
            "name": "fast-agent 升级 job",
            "parent_job_id": "",
            "sn": "PR2710R2014122206",
            "status": -2,
            "task_id": "746a4562c73cdde5fa1aec18dd980e21",
            "type_name": "agent_upgrade"
          }
        ]
      }

**如果输入空则为提示用户**

::

    $./cli job --filter-by-sn
    you should input --fiter-by-sn <sn>

**--filter-by-idc展示某个或者多个idc的状态**

::

    $./cli job --filter-by-idc shanghai12_ct
    
    ...
    ...

**如果输入空则为提示用户**
      
::    
      
    $./cli job --filter-by-idc
    you should input --fiter-by-sn <idc>

**--filter-by-job-id展示某个id的状态**

::

   $./cli job --filter-by-job-id 6e3fa09ba02483eb40bcf949bf44c157
   {
     "code": 0,
     "data": {
       "data": "{\"md5sum\":\"6582351ca8d6ac866713afeed5cc9827\",\"post_scripts\":\"-\",\"pre_scripts\":\"-\",\"timeout\":\"10\",\"url\":\"http://10.221.32.13/download/fast/download/fast/fast-agent/fast-agent_1.0.tar.gz\",\"version\":\"1.0\"}",
       "job_id": "6e3fa09ba02483eb40bcf949bf44c157",
       "job_template_id": 5,
       "name": "fast-agent 升级 job",
       "parent_job_id": "",
       "sn": "PR2710R2014122206",
       "status": -2,
       "task_id": "746a4562c73cdde5fa1aec18dd980e21",
       "type_name": "agent_upgrade"
     }
   }

**如果输入空则为提示用户**

::                        
             
    $./cli job --filter-by-job-id
    you should input --fiter-by-job-id <id>

子命令task使用说明
~~~~~~~~~~~~~~~~~~

**使用-h,--help获取帮助**

::

    $./cli task -h
    $./cli task --help
    $./cli task
    Do task operation such as task creating, status query

    Options:

      -h, --help
          display help information

      --status
          list all task working status

      --filter-by-sn
          list status by SN

      --filter-by-idc
          list status by IDC

      --filter-by-task-id
          list status by TaskID

      --create
          create task by file


**--status展示所有status状态**

::

    $./cli task --status
    ...展示太多，忽略

    ...

**--filter-by-sn展示某个或者多个sn的状态**

::

    //TODO:

**如果输入空则为提示用户**

::

    $./cli task --filter-by-sn
    you should input --fiter-by-sn <sn>

**--filter-by-idc展示某个或者多个idc的状态**

::

    $./cli task --filter-by-idc shanghai12_ct
    
    ...
    ...

**如果输入空则为提示用户**
      
::    
      
    $./cli task --filter-by-idc
    you should input --fiter-by-sn <idc>

**--filter-by-task-id展示某个id的状态**

::

    $./cli task --filter-by-task-id 66261921f295f3883b635e0fd827857e
    {
      "code": 0,
      "data": {
        "create_at": "2017-06-07T16:19:15+08:00",
        "creator": "haohui",
        "failed": 4,
        "jobs": {},
        "name": "视频 CACHE 部署",
        "status": 1,
        "success": 0,
        "task_id": "66261921f295f3883b635e0fd827857e",
        "task_template_id": 32,
        "total": 4,
        "update_at": "2017-06-07T16:21:38+08:00"
      }
    }


**如果输入空则为提示用户**

::                        
             
    $./cli task --filter-by-task-id
    you should input --fiter-by-task-id <id>


**--create来创建任务**
::

    $./cli task --create file
    {
       "code": 0,
       "message":""
    }

**如果输入空则为提示用户**

::

    $./cli task --create
    you should input --create <jsonfile>

子命令job-template的使用说明
~~~~~~~~~~~~~~~~~~~~~~

**-h,--help使用说明**

::

    $./cli job-template -h
    $./cli job-template --help
    $./cli job-template
    Do jobtemplate operation

    Options:

      -h, --help
          display help information

      --list
          list all jobtemplate

      --show
          show the jobtemplate data in by job template id

      --create
          create job template

**--list展示所有的template**

::

    $./cli job-template --list

    ...
    ...

**--show 展示某个id的templaete**

::

    ./cli job-template --show 26
    {
      "code": 0,
      "data": {
        "create_at": "2017-06-07T14:14:50+08:00",
        "id": 26,
        "meta_data": {
          "input": [
            "timeout",
            "pre_scripts",
            "version"
          ],
          "output": [
            "src_dir",
            "dst_dir",
            "timeout",
            "pre_scripts",
            "post_scripts"
          ],
          "template": {
            "custom": "version",
            "dst_dir": "${.app_nginx_config_dir}/current",
            "post_scripts": "mkdir -p /data/nginx-video/logs;rm -rf /opt/soft/nginx/conf;ln -sf ${.app_nginx_config_dir}/current /opt/soft/nginx/conf;service nginx restart",
            "src_dir": "${.app_nginx_config_dir}/${.version}",
            "type_name": "symlink"
          }
        },
        "name": "nginx 配置文件启用 job",
        "type_name": "symlink"
      }
    }

**如果输入为空则提示用户**

::

    ./cli job-template --show
    you should input --show <job template id>

**--create创建job模版**

::
    
    $./cli job-template --create <jsofile>
    {
       "code": 0,
        "message":""
    }


**如果输入为空则提示用户**

::

    ./cli job-template --create
    you should input --create <jsonfile>


子命令task-template使用说明
~~~~~~~~~~~~~



**-h,--help使用说明**

::

    $./cli task-template -h
    $./cli task-template --help
    $./cli task-template
    Do tasktemplate operation

    Options:

      -h, --help
          display help information

      --list
          list all tasktemplate

      --show
          show the tasktemplate data in by task template id

      --create
          create task template

**--list展示所有的template**

::

    $./cli task-template --list

    ...
    ...

**--show 展示某个id的templaete**

::

    $./cli task-template --show 28
    {
      "code": 0,
      "data": {
        "id": 28,
        "name": "fast-agent 升级任务",
        "topology": {
          "5": []
        }
      }
    }

**如果输入为空则提示用户**

::

    ./cli task-template --show
    you should input --show <task template id>

**--create创建job模版**

::
    
    $./cli task-template --create <jsofile>
    {
       "code": 0,
        "message":""
    }


**如果输入为空则提示用户**

::

    $./cli task-template --create
    you should input --create <jsonfile>


