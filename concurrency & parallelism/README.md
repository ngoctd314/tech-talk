# Concurrency & parallelism

## An Introduction to Concurrency

When most people use the word "concurrent", they're usually referring to a process that occurs simultaneously with one or more processes. It is also usually implied that all of these processes are making progress at about the same time. (stupid)

## Moore's Law, Web Scale, and the Mess We're In

**Moore's Law:** The number of components on an integrated circuit would double every two years.

Concurrency is so hard to get right.

## Why is concurrency Hard?

It's not uncommon for bugs to exist in code for years before some change in timing

**Critical section**

In concurrent programming, concurrent accesses to shared resources can lead to unexpected or erroneous behavior, so parts of the program where the shared resource is accesses need to be protected in ways that avoid the concurrent access. One way to do so is known as a critical section or critical region. This protected section cannot be entered by more than one process or thread at a time; others are suspended until the first leaves the critical section.

**Race Conditions**

A race condition occurs when two or more operations must execute in the correct order, but the program has not been written so that this order is guaranteed to be maintained.

Most of the time, this shows up in what's called a data race, where one concurrent operation attempts to read a variable while at some undetermined time another concurrent operation is attempting to write to the same variable.

```go
var data int
go func(){ //
    data++ // 3
}()
if data == 0 { // 5
    fmt.Printf("the value is %v.\n", data) // 6
}
```

Here, lines 3 and 5 are both trying to access the variable data, but there is no guarantee what order this might happen in. There are three possible outcomes to running this code:

- Nothing is printed. In this case, line 3 was executed before line 5
- "the value is 0." is printed. In this case lines 5 and 6  were executed before line 3.
- "the value is 1." is printed. In this case line 5 was executed before line 3, but line 3 was executed before line 6.

## Atomicity

When something is considered atomic, or to have the property of atomicity, this means that within the context that it is operating, it is indivisible, or uninterruptible.

The atomicity of an operation can change depending on the currently defined scope.

Example:

```go
i++;
```

This is about an simple an example as anyone can contrive, and yet it easily demonstrates the concept of atomicity. It may look atomic, but a brief analysis reveals several operations:

- Retrieve the value of i.
- Increment the value of i.
- Store the value of i.

While each of these operations alone is atomic, the combination of the three may not be, depending on your context. This reveals an interesting property of atomic operations: combining them does not necessarily produce a larger atomic operation. Making the operation atomic is dependent on which context you'd like it to be atomic within. If your context is a program with no concurrent processes, then this code is atomic within that context. If your context is a goroutine that doesn't expose i to other goroutines, then this code is atomic.

**So why we care?** Atomicity is important because if something is atomic, implicity it is safe within concurrent contexts?

If atomicity is the key to composing logically correct programs, and most statements aren't atomic, how do we reconcile these two statements? We have various techniques. The art then becomes determining which ares of your code need to be atomic.

## Memory Access Synchronization

Let's say we have a data race: two concurrent processes are attempting to access the same are of memory, and the way they are accessing the memory is not atomic.

```go
var data int
go func(){
    data++
}()
if data == 0 {
    fmt.Println("the value is 0.")
} else {
    fmt.Printf("the value is %v.\n", data)
}
```

In above example, we have three critical sections:

- Our goroutine, which is incrementing the data variables.
- Our if statement, which checks the value of data is 0.
- Our fmt.Printf statement, which retrievesd the value of data for output.

The one way to guard your program's critical section is to synchronize access to the memory between your critical sections.

```go
var memoryAccess sync.Mutex
var value int

go func(){
    memoryAccess.Lock()
    value++
    memoryAccess.Unlock()
}()

memoryAccess.Lock()
if value == 0 {
    fmt.Printf("the value is %v.\n", value)
} 
memoryAccess.Unlock()
```

We have solved our data race, we haven't actually solved our race condition! The order of operations in this program is still nondeterministic; we've just narrowed the scope of the nondeterminism a bit.ques of modeling concurrent problems, and we’ll discuss those in the next section.

## Deadlocks, Livelocks, and Starvation

**Deadlock**

A deadlocked program is one in which all concurrent processes are waiting on one another. In this state, the program will never recover without outside intervention.

