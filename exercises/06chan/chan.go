package channel

// this one is a bit more complex
// first, create a "delayedAdd" function that returns the sum of two ints but only after having waited for a few seconds
// then, call this function in a loop and see how we can't have the results if we don't wait for a very long time
// next, create a channel and run delayedAdd as a goroutine and make it respond on this channel
// when you call this new version of delayedAdd in a loop, you should get all the results much faster
