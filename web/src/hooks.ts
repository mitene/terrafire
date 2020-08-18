import {DependencyList, useCallback, useEffect, useRef, useState} from "react";

export function useAsync<T>(fn: () => Promise<T | undefined>, deps: DependencyList = []): [T | undefined, () => void] {
    const isMounted = useRef<boolean>(false);
    const [state, set] = useState<T | undefined>(undefined);
    const cb = useCallback(fn, deps);
    const callback = useCallback(
        () => {
            cb().then(v => isMounted.current && set(v)).catch(console.log);
        },
        [cb],
    );

    useEffect(() => {
        isMounted.current = true;

        callback();

        return () => { isMounted.current = false; }
    }, [callback]);

    return [state, callback]
}
