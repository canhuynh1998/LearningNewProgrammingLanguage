# Learning Golang Notes

<details open>
    <summary>Good Practices</summary>
        <blockquote>

### Formatting: use tab for indentation
### Naming conventions:
- Use mixedCaps or MixedCaps
        <blockquote>
</details>  

<details open>
        <summary>Pointer Stuff</summary>
            <blockquote>

```go
// pointer to integer
var ptr *int 

// get address of a variable AKA "point to"
// i is an int and ptr1 is a pointer to int
i1 := 1
ptr1 := &i 

// dereference a pointer 
// AKA get the value of the variable that the pointer points to
// The print statement will print out 2
i2 := 2
ptr2 := &i2
fmt.Println(*ptr2) 
```
</blockquote>

- Passing pointer around actually slowers than passing values due to **Escape Analysis**
    - **Escape Analysis** is basically checking whether the variable is in the heap or stack. If it is in the heap, it will need to be garbage collected -> this takes time. 
    - If variable is storeed in the stack, push/pop is sufficient and fast
- Pointers are good in following cases
    - Copying large structs
    - ??? *Mutating a variable when you pass into a function. By default, function is passing-by-value* ???
    - **If already used, then should keep using it for the API consistency**
    - To by-pass the default value. Example: int has default value 0 but pointer will be nil

</details>