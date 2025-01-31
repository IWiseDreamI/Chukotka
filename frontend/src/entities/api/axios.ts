import axios from 'axios';

const getConfig = async () => {
  return axios.get('/assets/config.json').then((res: any) => res.data);
};

export const config = await getConfig();

export const api = axios.create({
  baseURL: config.apiBaseUrl,
  headers: { 
    'content-type': 'application/json', 
  },
  timeout: 5000
});