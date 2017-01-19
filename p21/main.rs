fn q(n: usize, a: &[i32]) -> i32 {
    let mut max = 0;
    for i in 0..n {
        for j in i + 1..n {
            for k in j + 1..n {
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

fn q_alt(mut v: Vec<i32>) -> i32 {
    v.sort_by(|a, b| b.cmp(a));
    v[2..]
        .iter()
        .fold((v[0], v[1], 0), |(v1, v2, sum), x| {
            if v1 > v2 + x {
                return (v2, *x, 0);
            }
            if sum < v1 + v2 + x {
                return (v2, *x, v1 + v2 + x);
            }
            (v2, *x, sum)
        })
        .2
}

fn main() {
    let n: usize = 5;
    let a = [2, 3, 4, 5, 10];
    println!("{}", q(n, &a));

    let v: Vec<i32> = vec![2, 3, 4, 5, 10];
    println!("{}", q_alt(v));
}
