#include <stdio.h>

#define SIZE 10

int arr[SIZE] = {4, 2, 7, 5, 3, 8, 1, 9, 10, 6};
int temp[SIZE];

void mergeSort(int left, int right) {
    if (left >= right) return;
    int mid = (left + right) / 2;
    mergeSort(left, mid);
    mergeSort(mid + 1, right);
    
    int i = left, j = mid + 1, k = left;
    while (i <= mid && j <= right)
        temp[k++] = (arr[i] <= arr[j]) ? arr[i++] : arr[j++];
    while (i <= mid) temp[k++] = arr[i++];
    while (j <= right) temp[k++] = arr[j++];
    for (i = left; i <= right; i++) arr[i] = temp[i];
}

int main() {
    printf("정렬 전 배열: \n");
    for (int i = 0; i < SIZE; i++) printf("%d ", arr[i]);
    printf("\n");
    
    mergeSort(0, SIZE - 1);
    
    printf("정렬 후 배열: \n");
    for (int i = 0; i < SIZE; i++) printf("%d ", arr[i]);
    printf("\n");
    
    return 0;
}
