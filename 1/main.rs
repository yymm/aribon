fn q(n: usize, a: &[i32]) -> i32 {
    let mut max = 0;
    for i in 0..n {
        for j in i+1..n {
            for k in j+1..n {
                if a[i] + a[j] < a[k] || a[j] + a[k] < a[i] || a[k] + a[i] < a[j] {
                    continue;
                }
                let tmp = a[i] + a[j] + a[k];
                if max < tmp {
                    max = tmp;
                }
            }
        }
    }
    max
}
fn main(){
    let n: usize = 5;
    let a = [2, 3, 4, 5, 10];
    println!("{}", q(n, &a));
}
