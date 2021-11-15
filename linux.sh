# nginx日志 获取IP前k
awk '{print $1}' access.log | sort | uniq -c | sort -nr | head -n 10
awk '{aaa[$1]++;} END{for(i in aaa) { printf("%s\t%s\n", aaa[i], i); }}' test.log | sort -bn


 # 每行多个数，按列展示
 # eg.
 # 1 2 3
 # 4 5 6
 # res.
 # 1 2 3 4 5 6
 #
cat access.log | awk '{for (i=0;i<NF;i++){print $i;}}'


