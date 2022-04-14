# -*- coding: utf-8 -*-
# @Time    : 2019-09-05 14:19
# @Author  : Rick
# @File    : man.py.py
# @Software: PyCharm
import hashlib
import  pandas
import requests

from multiprocessing import Pool

class PhoneImeiConvert(object):

    def __init__(self):
        key = "dd852e8890ad4a85ab35a7a70111ca81"
        secret = "74ae2a77e1bc4b86b3e63c7f585967e0"
        self.auth = (key, secret)
        self.phone_imei_url = "http://nfjd-dp2apigw.jpushoa.com/x_v2/api_mapping/device/phn_num_hex_phn_imei"

    def get_imei_by_phone(self, phone_md5):
        data = {
            "data": phone_md5
        }
        response = requests.get(self.phone_imei_url, params=data, auth=self.auth)  # headers=self.headers)
        if response.status_code != 200:
            return None

        res = response.json()
        if res['code'] != 2000:
            return None

        imei = res['data'].pop(0) if 'data' in res and res['data'] else None
        return imei

    def md5(self, raw_str):
        return hashlib.new("md5", raw_str.encode("utf8")).hexdigest()

l1 = []
with open("part-00000-921d6fb2-4c18-4401-b006-0a71e61da147-c000.csv") as fp:
    for line in fp.readlines():
        l1.append(line.replace("\n", ""))


print(len(l1))
def run(filename):
    count = 1
    with open(filename,"a+") as wp:
        for x in l1:
            count += 1
            try:
                imei = PhoneImeiConvert().get_imei_by_phone(x)
            except:
                imei = None
            _str = "{},{} \n".format(x, imei)
            print(count, _str)
            wp.writelines(_str)
print(len(l1))
def run2(x):
    count = 1
    with open("zuoak.csv","a+") as wp:
        count += 1
        try:
            imei = PhoneImeiConvert().get_imei_by_phone(x)
        except:
            imei = None
        _str = "{},{} \n".format(x, imei)
        print(count, _str)
        wp.writelines(_str)
# p = Pool()
# for i in range(8):
#     filename = "result_{}.csv".format(i)
#     p.apply_async(run, args=(filename,))
#
# print('Waiting for all subprocesses done...')
# p.close()
# p.join()
# print('All subprocesses done.')
from multiprocessing.dummy import Pool as ThreadPool
pool = ThreadPool()
results = pool.map(run2, l1)
pool.close()
pool.join()

# cat result_* > merge.csv

#imei = PhoneImeiConvert().get_imei_by_phone("12ce69ef3d957cbf35740a821f7f9ff9")
#print(imei)