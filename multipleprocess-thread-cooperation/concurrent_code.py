# -*- coding: utf-8 -*-
# @Time    : 2019/10/22 7:21 下午
# @Author  : Rick
# @File    : concurrent_code.py
# @Software: PyCharm

'''
python 并发
1.进程+协程
2.多线程并发
'''
import os
import time
from datetime import datetime

def print_time():

    print(datetime.now())

def info(title):
    print(title)
    print('module name:', __name__)
    print('parent process:', os.getppid())
    print('process id:', os.getpid())

# -*- coding=utf-8 -*-
import requests
from multiprocessing import Process

def foo(name):
    info('main line')
    print('hello', name)

if __name__ == '__main__':
    p = Process(target=foo, args=('bob',))
    p.start()
    p.join()

#
# def process_start(url_list):
#     tasks = []
#     for url in url_list:
#         tasks.append(gevent.spawn(fetch, url))
#     gevent.joinall(tasks)  # 使用协程来执行
#
#
# def task_start(filepath, flag=100000):  # 每10W条url启动一个进程
#     url = ""
#     url_list = []  # 这个list用于存放协程任务
#     i = 0  # 计数器，记录添加了多少个url到协程队列
#     while url != '':
#         i += 1
#         url_list.append(url)  # 每次读取出url，将url添加到队列
#         if i == flag:  # 一定数量的url就启动一个进程并执行
#             p = Process(target=process_start, args=(url_list,))
#             p.start()
#             url_list = []  # 重置url队列
#             i = 0  # 重置计数器
#
#     if url_list :  # 若退出循环后任务队列里还有url剩余
#         p = Process(target=process_start, args=(url_list,))  # 把剩余的url全都放到最后这个进程来执行
#         p.start()
#
#
# if __name__ == '__main__':
#     task_start('./testData.txt')  # 读取指定文件
