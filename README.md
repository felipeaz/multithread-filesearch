# multithread-filesearch
Multithread file searching on Mac OS

Simply code that reads all files using recursive function and go routines and stores all matched
filepath which has the README.md name.

This example uses mutex to lock the shared variable ``matches`` to avoid race condition and uses 
``wait group`` to synchronize the goroutines