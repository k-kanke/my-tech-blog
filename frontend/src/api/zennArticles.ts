import axios from 'axios';

const API_BASE_URL = 'http://localhost:8080';

export const fetchZennArticles = () =>
  axios.get(`${API_BASE_URL}/zenn_articles`);
