# -*- coding: utf-8 -*-

import tensorflow as tf
import numpy as np
import os
import warnings

warnings.filterwarnings("ignore")
os.environ['TF_CPP_MIN_LOG_LEVEL'] = '3'


def _tensordot_test():
    a = tf.ones(shape=[2, 2, 3])
    b = tf.ones(shape=[3, 2, 6])
    # c = tf.tensordot(a, b, axes=((1, 2), (0, 1)))
    # c = tf.tensordot(a, b, axes=1)
    c = tf.tensordot(a, b, axes=2)
    with tf.compat.v1.Session() as sess:
        sess.run(tf.compat.v1.global_variables_initializer())
        print(sess.run(c))
        print(sess.run(tf.shape(c)))


def _tf_hl_test():
    hello = tf.constant("Hello, TensorFlow!")
    with tf.compat.v1.Session() as sess:
        print(sess.run(hello))


def _tensordot_np_test():
    a = np.arange(60.).reshape(3, 4, 5)
    b = np.arange(24.).reshape(4, 3, 2)
    c = np.tensordot(a, b, axes=([1, 0], [0, 1]))
    print(c.shape)

    print("-1" * 36)
    np.random.seed(10)
    A = np.random.randint(0, 9, (2, 5, 4))
    B = np.random.randint(0, 9, (4, 5, 3))
    C = np.tensordot(A, B, [(2, 1), (0, 1)])
    # C = np.tensordot(A, B, axes=2)
    # C = np.tensordot(A, B, axes=1)
    print(C.shape)
    print(C)

    print("-2" * 36)
    np.random.seed(10)
    A = np.random.randint(0, 9, (2, 4, 5))
    B = np.random.randint(0, 9, (4, 5, 3))
    C = np.tensordot(A, B, [(1, 2), (0, 1)])
    print(C.shape)
    print(C)

    print("-3" * 36)
    np.random.seed(10)
    A = np.random.randint(0, 9, (2, 4, 5))
    B = np.random.randint(0, 9, (4, 5))
    C = np.tensordot(A, B, [(1, 2), (0, 1)])
    print(C.shape)
    print(C)

    print("-4" * 36)
    np.random.seed(10)
    A = np.random.randint(0, 9, (4, 5))
    B = np.random.randint(0, 9, (4, 5, 3))
    C = np.tensordot(A, B, [(0, 1), (0, 1)])
    print(C.shape)
    print(C)

    print("-5" * 36)
    np.random.seed(10)
    A = np.random.randint(0, 9, (4, 5, 2))
    B = np.random.randint(0, 9, (4, 5, 3))
    C = np.tensordot(A, B, [(0, 1), (0, 1)])
    print(C.shape)
    print(C)

    print("-6" * 36)
    np.random.seed(10)
    A = np.random.randint(0, 9, (2, 4, 5))
    B = np.random.randint(0, 9, (4, 5, 3))
    C = np.tensordot(A, B, axes=2)
    # C = np.tensordot(A, B, axes=1)
    print(C.shape)
    print(C)

    print("-7" * 36)
    np.random.seed(10)
    A = np.random.randint(0, 9, (2, 3, 4))
    B = np.random.randint(0, 9, (4, 5, 3))
    C = np.tensordot(A, B, axes=1)
    print(C.shape)
    print(C)


if __name__ == '__main__':
    _tensordot_np_test()
