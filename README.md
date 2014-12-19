# tesserdata

Extracting statistics from a game of [tesseract](http://tesserract.gg).
Example data below is from a game between two players suggestively labelled 'arne' and 'mathias'.

## tesserstamp

Records frag time stamps form tesseract stdout. Usage:
``` 	tesseract | tesserstamp > timestamps.txt  ```
Output: time stamp in ms, boolean indication if you got fragged.


## tesserrate

Extracts fragging rates during game progress. E.g., here player A got a bit tired near the end of the game.

![fig](fig/frag_rate_A.png)
![fig](fig/frag_rate_B.png)

## tesserhist

Makes a histogram of time between successive frags. E.g., here player A most often fragged one second after the previous frag. Intense game.

![fig](fig/frag_hist_A.png)
![fig](fig/frag_hist_B.png)


## tesserlife

![fig](fig/lifetime_A.png)
![fig](fig/lifetime_B.png)

## tessercorr
![fig](fig/frag_autocorr_A.png)
![fig](fig/frag_autocorr_B.png)
![fig](fig/frag_corrAB.png)
![fig](fig/frag_corrBA.png)
