import { useState, useEffect } from 'react';
import axios from 'axios';

function App() {
  const [maxTime, setMaxTime] = useState(60)
  const [counter, setCounter] = useState(0);
  const [seconds, setSeconds] = useState<number>(maxTime);
  const [isRunning, setIsRunning] = useState<boolean>(false);
  const [images, setImages] = useState<string[]>([]);

  const resetTimer = () => {
    setIsRunning(true);
    setSeconds(maxTime);
  };

  function nextImage() {
    setCounter(counter + 1);
    resetTimer();

    console.log(images.length + " " + counter)
    if (counter < images.length) {
      return
    }

    const fetchImages = async () => {
      const response = await axios.get('http://localhost:8080/api/mydata/random', {
        responseType: 'blob', // or 'blob' if binary
      })
      
      const url = URL.createObjectURL(response.data); // 👈 Convert Blob to Object URL
      setImages(prev => [...prev, url]);
    };

    fetchImages();
  }

  function previousImage() {
    if (counter > 0) {
      setCounter(counter - 1);
      resetTimer();
    }
  }

  function togglePause() {
    setIsRunning(!isRunning)
  }

  useEffect(() => {
    let intervalId: NodeJS.Timeout | undefined;

    if (isRunning && seconds > 0) {
      intervalId = setInterval(() => {
        setSeconds((prevSeconds) => prevSeconds - 1);
      }, 1000);
    } else if (seconds === 0) {
      nextImage()
    }

    return () => {
      if (intervalId) {
        clearInterval(intervalId);
      }
    };
  }, [isRunning, seconds]);

  return (
    <div id="image-container">
      <button onClick={previousImage}>Previous Image</button>
      <button onClick={nextImage}>Next Image</button>
      <button onClick={togglePause}>{isRunning?"Pause":"Play"}</button>
      <input type="number" id="time" name="time" value={maxTime} onChange={e => setMaxTime(+e.target.value)}></input>

      <div id="counter">{"Images Shown: " + counter}</div>
      
      {images.length > 0 ? (
        <img src={images[counter - 1]} alt="Fetched from backend" />
      ) : (
        <p>Loading image...</p>
      )}
      {/* <img src={imageUrl} alt="image" style={{ maxWidth: '500px' }} /> */}
      
      <div>{"Time Remaining: " + seconds + "s"}</div>
      {/* <img id="image" src="" alt="Random Image" style="display:none;"> */}
    
    </div>
  );
}

export default App;