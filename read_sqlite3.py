# -*- coding:utf-8 -*-

import pandas
import csv, sqlite3

if __name__ == '__main__':
    conn= sqlite3.connect("app_info.sqlite")
    df = pandas.read_sql('select * from apps', conn)
    df.to_excel('apps.xlsx', index=False)
