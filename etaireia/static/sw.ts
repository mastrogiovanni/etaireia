import {precacheAndRoute} from 'workbox-precaching'
import {registerRoute} from 'workbox-routing'
import {NetworkFirst} from 'workbox-strategies'

declare const self: ServiceWorkerGlobalScope;

// precacheAndRoute(self.__WB_MANIFEST);

// registerRoute(
//     ({url}) => {
//         return url.origin === "https://www.cerchidonda.org" || url.pathname.startsWith("/api/v1")
//     }, 
//     new NetworkFirst({cacheName: 'images'})
// );
