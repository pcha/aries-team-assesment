import { useState, useEffect } from 'react';

/**
 * Width and Height dimensions
 */
type Dimensions = {
  width: number,
  height: number
}

/**
 * gets current window dimensions
 */
function getWindowDimensions(): Dimensions {
  const { innerWidth: width, innerHeight: height } = window;
  return {
    width,
    height
  };
}

/**
 * Hook to listen window resizing and return real time dimensions
 *
 * @return Window dimensions
 */
function useWindowDimensions(): Dimensions {
  const [windowDimensions, setWindowDimensions] = useState(getWindowDimensions());

  useEffect(() => {
    function handleResize() {
      setWindowDimensions(getWindowDimensions());
    }

    window.addEventListener('resize', handleResize);
    return () => window.removeEventListener('resize', handleResize);
  }, []);

  return windowDimensions;
}

export default useWindowDimensions