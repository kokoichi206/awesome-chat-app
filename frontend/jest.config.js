/** @type {import('ts-jest/dist/types').InitialOptionsTsJest} */
module.exports = {
  preset: 'ts-jest',
  testEnvironment: 'jest-environment-jsdom',
  moduleFileExtensions: ['js', 'ts', 'json', 'vue'],
  moduleNameMapper: {
    '^@/(.*)$': '<rootDir>/src/$1',
    'tests/(.*)$': '<rootDir>/tests/$1',
  },
  transform: {
    '^.+\\.vue$': 'vue-jest',
    '^.+.(js|jsx|ts)$': 'babel-jest',
  },
  collectCoverage: true,
  collectCoverageFrom: [
    './src/components/**/*.{js,ts}',
    './src/modules/**/*.{js,ts}',
    './src/store/**/*.{js,ts}',
    './src/views/**/*.{js,ts}',
  ],
  globals: {
    'ts-jest': {
      tsconfig: false,
      useESM: true,
      babelConfig: true,
      plugins: ['babel-plugin-transform-vite-meta-env'],
    },
  },
};