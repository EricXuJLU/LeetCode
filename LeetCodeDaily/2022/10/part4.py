class Solution:
    def countStudents(self, students: List[int], sandwiches: List[int]) -> int:
        cnt = [0, 0]
        for s in students:
            cnt[int(s)] += 1
        for x in sandwiches:
            cnt[int(x)] -= 1
            if cnt[x] < 0:
                return cnt[x]+cnt[1-x]+1
        return 0
