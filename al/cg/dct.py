import cv2
import numpy as np

np.set_printoptions(linewidth=400)


testdata = np.array([
        [61., 19., 50., 20.],
        [82., 26., 61., 45.],
        [89., 90., 82., 43.],
        [93., 59., 53., 97.],
    ])

def fdct(mat):
    w, h = np.size(mat, 1), np.size(mat, 0)
    A = np.zeros([h, w])

    for i in range(0, w):
        for j in range(0, h):
            if i == 0:
                a = np.sqrt(1/w)
            else:
                a = np.sqrt(2/w)
            A[i][j] = a*np.cos(np.pi*(j+0.5)*i/w)

    return np.around(A.dot(mat).dot(A.T), decimals=8)


def showmat(mat):
    for line in mat:
        print(line)


def is_equal(a, b):
    h, w = len(a), len(a[0])
    for i in range(0, h):
        for j in range(0, w):
            if a[i][j] != b[i][j]:
                print(a[i][j], b[i][j])
                return False
    return True

def test_standard_fdct(mat):
    print('--cv2---')
    standard = cv2.dct(mat)
    showmat(standard)

def test_mine_fdct(mat):
    print('--mine--')
    mine = fdct(mat)
    showmat(mine)

def test_mat_dot():
    a = np.array([
        [1., 2., 3.],
        [3., 4., 5.],
    ])
    b = np.array([
        [1., 2.],
        [3., 4.],
        [3., 4.],
    ])
    x = a.dot(b)
    showmat(a)
    print()
    showmat(b)
    print()
    showmat(x)


if __name__ == '__main__':
    # test_standard_fdct()
    # test_mine_fdct()
    # test_mat_dot()
    for i in range(0, 10000):
        fdct(testdata)
