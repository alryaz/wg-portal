import { authStore } from '../stores/auth';

export const fetchWrapper = {
    get: request('GET'),
    post: request('POST'),
    put: request('PUT'),
    delete: request('DELETE')
};

export const apiWrapper = {
    get: apiRequest('GET'),
    post: apiRequest('POST'),
    put: apiRequest('PUT'),
    delete: apiRequest('DELETE')
};

function request(method) {
    return (url, body = undefined) => {
        const requestOptions = {
            method,
            headers: authHeader(url)
        };
        if (body) {
            requestOptions.headers['Content-Type'] = 'application/json';
            requestOptions.body = JSON.stringify(body);
        }
        return fetch(url, requestOptions).then(handleResponse);
    }
}

function apiRequest(method) {
    return (path, body = undefined) => {
        const url = WGPORTAL_BACKEND_BASE_URL + "/frontend/api/v1" + path
        const requestOptions = {
            method,
            headers: authHeader(url)
        };
        if (body) {
            requestOptions.headers['Content-Type'] = 'application/json';
            requestOptions.body = JSON.stringify(body);
        }
        return fetch(url, requestOptions).then(handleResponse);
    }
}

// helper functions

function authHeader(url) {
    // return auth header with jwt if user is logged in and request is to the api url
    const auth = authStore();
    const isApiUrl = url.startsWith(WGPORTAL_BACKEND_BASE_URL);
    if (auth.IsAuthenticated && isApiUrl) {
        return { 'X-Frontend-User': `Identifier ${auth.UserIdentifier}` };
    } else {
        return {};
    }
}

function handleResponse(response) {
    return response.text().then(text => {
        const data = text && JSON.parse(text);

        if (!response.ok) {
            const auth = authStore();
            if ([401, 403].includes(response.status) && auth.IsAuthenticated) {
                // auto logout if 401 Unauthorized or 403 Forbidden response returned from api
                auth.Logout();
            }

            const error = (data && data.message) || response.statusText;
            return Promise.reject(error);
        }

        return data;
    });
}