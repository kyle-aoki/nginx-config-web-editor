import { useEffect } from "react";

export const useAsyncEffect = (effect: () => Promise<void>, deps?: React.DependencyList | undefined): void => {
  useEffect(() => {
    effect();
  }, deps);
};
