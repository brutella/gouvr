/*
Defines types and implements functionality to read values for an UVR 1611 data bus.

The UVR1611 sends 64 bytes packets preceded and folled by 16 high bits to synchronize transmissions.
Some functionality in this package can be reused for other UVR devices; but make sure to read the
documentation of the data bus first.
*/
package uvr1611