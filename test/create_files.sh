#!/bin/zsh

touch ./files/test_1old.txt
setfile -d "$(date -v-1d +'%m/%d/%Y %H:%M:%S')" ./files/test_1old.txt
setfile -m "$(date -v-1d +'%m/%d/%Y %H:%M:%S')" ./files/test_1old.txt

touch ./files/test_7old.txt
setfile -d "$(date -v-7d +'%m/%d/%Y %H:%M:%S')" ./files/test_7old.txt
setfile -m "$(date -v-7d +'%m/%d/%Y %H:%M:%S')" ./files/test_7old.txt

touch ./files/test_10old.txt
setfile -d "$(date -v-10d +'%m/%d/%Y %H:%M:%S')" ./files/test_10old.txt
setfile -m "$(date -v-10d +'%m/%d/%Y %H:%M:%S')" ./files/test_10old.txt

touch ./files/test_30old.txt
setfile -d "$(date -v-30d +'%m/%d/%Y %H:%M:%S')" ./files/test_30old.txt
setfile -m "$(date -v-30d +'%m/%d/%Y %H:%M:%S')" ./files/test_30old.txt

touch ./files/test_365old.txt
setfile -d "$(date -v-365d +'%m/%d/%Y %H:%M:%S')" ./files/test_365old.txt
setfile -m "$(date -v-365d +'%m/%d/%Y %H:%M:%S')" ./files/test_365old.txt
