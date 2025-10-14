package main

import sort


func answerQueriesWithBS(nums []int, queries []int) []int {
    sort.Ints(nums)
    prefixes := make([]int,len(nums))

    ans := make([]int,len(queries))
    prefixes[0] = nums[0]
    for i:=1; i < len(nums); i++{
        prefixes[i] = prefixes[i-1] + nums[i]
    } 

    for qi, q := range queries{
        ans[qi] = findFirstLessBS(prefixes,q) +1
    }

    return ans
}

func findFirstLessBS(prefixes []int,target int) int{
    l,r:= 0 , len(prefixes)
    for l<r{
        t:= l + (r-l)/2
        if prefixes[t] <= target{
            l = t+1
        }else{
            r = t
        }
    }


    return l-1
}



func answerQueriesTwoSortTwoPointers(nums []int, queries []int) []int {
    sort.Ints(nums)
    prefixes := make([]int,len(nums))

    ans := make([]int,len(queries))
    prefixes[0] = nums[0]
    for i:=1; i < len(nums); i++{
        prefixes[i] = prefixes[i-1] + nums[i]
    } 

    type query struct{
        val int
        pos int
    }

    nq := make([]query,0,len(nums))
    for i,q:= range queries{
        nq = append(nq,query{val:q,pos:i})
    }
    sort.Slice(nq, func(i, j int) bool {
        return nq[i].val < nq[j].val
    })

    pi:=0
    for _,q := range nq {
        for pi < len(prefixes) && prefixes[pi] <= q.val {
            pi++
        }
        ans[q.pos] = pi
    }

    return ans
}