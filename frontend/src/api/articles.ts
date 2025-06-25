import axios from "axios";

const API_BASE_URL = 'http://localhost:8080';

export const fetchArticles = () => 
    axios.get(`${API_BASE_URL}/articles`)