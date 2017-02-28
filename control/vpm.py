# -*- coding: utf-8 -*-

import os
import sys
import linecache

lastsearch = "/Users/CodeMania/.vpm/lastsearch.dat"
rootpath = "/Users/CodeMania/Documents/编程相关"

def openf(args):
    filenum = int(args[0])
    fileinfo = linecache.getline(lastsearch, filenum)
    filepath = rootpath + fileinfo.split(' ')[5][1:]
    os.system("open " + filepath)
    return

def search(args):
    keyword = args[0]
    filetype = args[1]
    os.system("~/.vpm/search" + " " + keyword + " " + filetype)
    return

def update(args):
    os.system("~/.vpm/engine")
    return

commandtable = {
    "-o": openf,
    "-u": update,
    "-s": search
}

if __name__ == "__main__":
    # if len(sys.argv) > 2:
    #     args = sys.argv[2:]
    # else:
    #     args = []
    #
    # commandtable[sys.argv[1]](args)
    commandtable["-s"](["java", "pdf"])









