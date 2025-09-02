package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// Пишите тесты в этом файле
func TestGenerateRandomElements(t *testing.T){
	testingDataNegativ:=[]int{-100_000_000,-1000,-9,-8,-7,-1,0}
	for _,j:=range testingDataNegativ {
		result := generateRandomElements(j)
		require.Len(t,result,0, "при тесте были заданы нулевые или отрицательные значения массива, тест не пройден")
	}
testingDataPozitiv:=[]int{1,7,8,9,100,300,500,1000,5000,20000,100000000}

for _,j:=range testingDataPozitiv {
result := generateRandomElements(j)	
		require.Len(t,result,j,fmt.Sprintf("Ожидаемая длина: %d, фактическая: %d", j,result))
	}
// расчет среднего значения 
for _,j:=range testingDataPozitiv {
	sum:=0

for _, value := range generateRandomElements(j){
	sum +=value
	}	
	average:=float64(sum)/float64(j) 

if average>float64(j)||average<0 {
        t.Errorf("ожидаемое значение вне пределов допустимых значений  минимум %d максимум %d, получено %.2f", 0, j, average)
    }	
if j>500 {	
minValue := average*0.9
maxValue := average*1.1
    if average < minValue || average > maxValue {
        t.Errorf("ожидаемое значение вне пределов ожидаемого распределения от%.2f до%.2f, получено %.2f", minValue, maxValue, average)
    }	}}}
	func TestMaximum(t *testing.T){
data:=[]int{0,1,2,3,4,5,6,7,8,9,10,100,200,300,500,1000,10000,100000,100000000}
require.Equal(t,maximum(data),100000000,"тест с положительными числами не пройден")
dataNegative:=[]int{-1,-2,-3,-4,-5,-6,-7,-8,-9,-10,-100,-200,-300,-500,-1000,-10000,-100000,-100000000}
require.Equal(t,maximum(dataNegative),-1,"тест с отрицательными числами не пройден")
dataMix:=[]int{0,-1,2,-3,4,-5,6,-7,8,-9,10,100,-200,300,-500,1000,-10000,100000,-100000000}
require.Equal(t,maximum(dataMix),100000,"тест со смешанными числами не пройден")
dataEqual:=[]int{20,20,20,20,20,20,20}
require.Equal(t,maximum(dataEqual),20,"тест с одинаковыми числами не пройден")
datasingl:=[]int{222}
require.Equal(t,maximum(datasingl),222,"тест с одним числом не пройден")
//тест на панику при пустом срезе
       defer func() {
        r := recover()
			require.Equal(t, "Ошибка: пустой срез", r, "тест с пустым срезом не пройден")
            require.NotNil(t, r, "Ожидалась паника при пустом срезе")
        }()
        maximum([]int{})    
	}          
   func TestComparisonMax(t *testing.T) {
    // Генерируем тестовый массив
    testData := generateRandomElements(100000000)
    max1 := maximum(testData)
    max2 := maxChunks(testData)
    require.Equal(t, max1, max2, "Результаты поиска максимума отличаются. Ошибка")
}    
        