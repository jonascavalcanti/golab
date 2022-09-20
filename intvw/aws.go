/*
Find Non-overlapping intervals among a given set of intervals

Given **N** set of time intervals, the task is to find the intervals which don’t overlap with the given set of intervals.

Example 1:
    input: [1, 3], [2, 4], [3, 5], [7, 9]
    output: [5, 7]
    
Example 2:
    input: [1, 3], [9, 12], [2, 4], [6, 8]
    output: [4, 6], [8, 9]

1  2  3  4  5  6  7 8 9
[1, 3][3, 5]      [7, 9]
   [2, 4]    ^
    
   [2, 4]
[1, 3], [2, 4], [6, 8] [9, 12]

*/



import main

func main (){
    
    map:=make(map[int]int)
    map[0]=[1, 3]
    map[1]=[2, 4]
    map[2]= [3, 5]
    map[3]=[7, 9]
    
    for key,value := range map{
        for i:=key; i < len(map);i++{
            
        }
    }
    
}