# WordCount_Destributed

- Input: a file containing English text , as in “ExampleIn.txt”
- Output: a  file  containing  each  unique  word  and  its  associated  count  as  appeared  in  the  input  text . The  output  file format should follow the “ExampleOut.txt” provided. “ExampleOut.txt”.

InputFileName: test.txt , OutputFileName: WordCountOutput.txt

- Method: First Convert Input to lower case so the output follows “ExampleOut.txt” format.
The input file should be divided evenly among 5 go routines, each routine computes the word counts for the portion of  the  file  it  is  responsible  for.  
After  each  routine finishes,  it  writes  the  output  to  a  shared  map  that  is  handled  by another routine “reducer”. e.g. if one routine has the word “assignment” 3 times and another has the same word 7 times then the map entry for the word assignment should now be 10.

When all routines write the output to the shared map, the “reducer” should write the output in the file sorted by the frequency (sorting is not distributed, “reducer” can be responsible for it).If two or more words have the same frequency, they ae sorted alphabetically.
