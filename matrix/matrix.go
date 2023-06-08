package matrix

import (
	"bytes"
	"fmt"

	"github.com/syuya2036/num/functions"
)


// GaussSeidelは、係数行列matと定数ベクトルyを受け取り、
// ガウス・ザイデル法を用いて連立方程式を解きます。
// 初期値はすべて0で、収束条件は差分の絶対値がeps以下になることです。
// 返り値は解ベクトルxです。
func GaussSeidel(mat, y Matrix) Matrix {
	x := Zeros(1, mat.Cols) // 初期値はすべて0
	eps := 1.0e-6 // 収束条件は差分の絶対値がeps以下

	for n := 0; n < 3000000; n++ { // 最大で3000000回までループ
		xBefore := x.Copy() // 前回のxを保存
		diff := 0.0 // 差分の絶対値の合計

		for i := 0; i < mat.Rows; i++ {
			sum := y.Get(0,i) // 初期値は定数ベクトルyのi番目の要素
			for j := 0; j < mat.Cols; j++ {
				if i == j {
					continue
				}
				sum -= mat.Get(i,j) * x.Get(0,j)
			}
			x.Asign(0,i,sum/mat.Get(i,i))

			diff += functions.Abs(x.Get(0,i) - xBefore.Get(0,i)) // 差分の絶対値を加算
		}

		if diff < eps { // 収束条件を満たしたら終了
			break
		}
	}
	
	return x
}


func (mat Matrix) GaussianElimination() Matrix {

	for i := 0; i < mat.Rows; i++ {
		mat = mat.MulScalerToRow(i, 1/mat.Get(i, i))

		for j := 0; j < mat.Rows; j++ {
			if i == j {
				continue
			}

			// 引く側の行にj行i列の値をかける
			mat = mat.SubRowToRow(j, i, mat.Get(j, i))
		}
	}
	return mat
}

type Matrix struct {
	Mat  []float64
	Rows int
	Cols int
}

// 二次元配列から初期化
func NewMatrixFromArray(mat [][]float64) Matrix {

	rows, cols := len(mat), len(mat[0])

	matC := make([]float64, rows*cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			matC[i*cols+j] = mat[i][j]
		}
	}

	return NewMatrix(matC, rows, cols)
}

func NewMatrix(mat []float64, rows, cols int) Matrix {
	var newMat *Matrix = new(Matrix)

	newMat.Mat = mat
	newMat.Rows, newMat.Cols = rows, cols

	return *newMat
}

// 多次元配列ライクにアクセス
func (mat Matrix) Get(i, j int) float64 {
	return mat.Mat[i*mat.Cols+j]
}

// 指定した行に値をかける
func (mat Matrix) MulScalerToRow(idx int, scaler float64) Matrix {
	newMat := make([]float64, mat.Cols*mat.Rows)
	copy(newMat, mat.Mat)
	
	for i := 0; i < mat.Cols; i++ {
		newMat[idx*mat.Cols+i] *= scaler
	}

	return NewMatrix(newMat, mat.Rows, mat.Cols)
}

// 指定した行から指定した行を引く
func (mat Matrix) SubRowToRow(idx1, idx2 int, scaler float64) Matrix {
	newMat := make([]float64, mat.Cols*mat.Rows)
	copy(newMat, mat.Mat)
	
	for i := 0; i < mat.Cols; i++ {
		newMat[idx1*mat.Cols+i] -= mat.Mat[idx2*mat.Cols+i] * scaler
	}

	return NewMatrix(newMat, mat.Rows, mat.Cols)
}

//　行列をそれっぽく表示
func (mat Matrix) String() string {
	var matString bytes.Buffer

	matString.WriteString("--------------begin---------------\n")
	for i := 0; i < mat.Rows; i++ {
		for j := 0; j < mat.Cols; j++ {
			matString.WriteString(fmt.Sprintf("%f, ", mat.Mat[i*mat.Cols+j]))
		}
		matString.WriteString("\n")
	}

	matString.WriteString("---------------end----------------")
	return matString.String()
}

func (mat Matrix) Asign(i, j int, scaler float64) {
	mat.Mat[i*mat.Cols+j] = scaler
}

func Zeros(rows, cols int) Matrix {
	mat := make([]float64, rows*cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			mat[i*cols+j] = 0
		}
	}

	return NewMatrix(mat, rows, cols)
}

func (mat Matrix) Copy() Matrix {
	newMat := make([]float64, mat.Cols*mat.Rows)
	copy(newMat, mat.Mat)
	return NewMatrix(newMat, mat.Rows, mat.Cols)
}