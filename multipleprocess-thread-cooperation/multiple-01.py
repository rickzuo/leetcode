# -*- coding: utf-8 -*-
# @Time    : 2019/10/23 10:39 上午
# @Author  : Rick
# @File    : multiple-01.py
# @Software: PyCharm
import multiprocessing
import time
import requests
from bs4 import BeautifulSoup

def request_douban(url):
   try:
       response = requests.get(url)
       if response.status_code == 200:
           return response.text
   except requests.RequestException:
       return None

def save_content(soup):
    list = soup.find(class_='grid_view').find_all('li')
    for item in list:
        try:
            item_name = item.find(class_='title').string
            item_img = item.find('a').find('img').get('src')
            item_index = item.find(class_='').string
            item_score = item.find(class_='rating_num').string
            item_author = item.find('p').text
            item_intr = item.find(class_='inq').string

            # print('爬取电影：' + item_index + ' | ' + item_name +' | ' + item_img +' | ' + item_score +' | ' + item_author +' | ' + item_intr )
            print('爬取电影：' + item_index + ' | ' + item_name + ' | ' + item_score + ' | ' + item_intr)
        except:
            continue

def save_to_xml(data):
    import xlwt

    book = xlwt.Workbook(encoding='utf-8', style_compression=0)

    sheet = book.add_sheet('豆瓣电影Top250', cell_overwrite_ok=True)
    sheet.write(0, 0, '名称')
    sheet.write(0, 1, '图片')
    sheet.write(0, 2, '排名')
    sheet.write(0, 3, '评分')
    sheet.write(0, 4, '作者')
    sheet.write(0, 5, '简介')

    # sheet.write(n, 0, item_name)
    # sheet.write(n, 1, item_img)
    # sheet.write(n, 2, item_index)
    # sheet.write(n, 3, item_score)
    # sheet.write(n, 4, item_author)
    # sheet.write(n, 5, item_intr)

    book.save(u'豆瓣最受欢迎的250部电影.xlsx')

def run(url):
    html = request_douban(url)
    soup = BeautifulSoup(html, 'lxml')
    save_content(soup)


if __name__ == '__main__':
    start = time.time()
    urls = []
    ct = multiprocessing.cpu_count()
    print(ct)
    pool = multiprocessing.Pool(ct)
    for i in range(0, 100):
        url = 'https://movie.douban.com/top250?start=' + str(i * 25) + '&filter='
        urls.append(url)

    pool.map(run, urls)
    pool.close()
    pool.join()