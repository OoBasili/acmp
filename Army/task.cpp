#include <iostream>
#include <vector>
 
struct Set {
    std::vector<int> a;
    Set(int n): a(n, 0) {};
    int d(int k) {
        int s = 0;
        for (int i = k; i >= 0; i = (i & (i + 1)) - 1) {
            s += a[i];
        }
        return s;
    }
    void add(int v) {
        while (v < (int)a.size()) {
            a[v]++;
            v = v | (v + 1);
        }
    }
};
   
int main()
{
    int n, m, s=0;
    std::cin >> n >> m;
    for (int i = 0; i < m; i++)
    {
        Set se(n);
        for (int j = 0; j < n; j++)
        {
            int v;
            std::cin >> v;
            v--;
            s += se.d(n - 1) - se.d(v);
            se.add(v);
        }
    }
    std::cout << s;
}