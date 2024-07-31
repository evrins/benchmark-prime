#!/bin/bash

myArray=(1000000 2000000 3000000 4000000 5000000 6000000 7000000 8000000 9000000
  10000000 20000000 30000000 40000000 50000000 60000000 40000000 40000000 40000000 40000000
)

seq1=$(seq 1000000 1000000 9000000 | xargs printf "%.0f\n")
seq2=$(seq 10000000 10000000 100000000 | xargs printf "%.0f\n")

combined_seq="$seq1 $seq2"

echo 'lang, total, count, time'

for i in $combined_seq ; do
#  ./go/prime $i
#  java -jar java/app/build/libs/app.jar $i
#  ./rust/target/release/prime $i
#  python py/main.py $i
  node nodejs/index.js $i
done
