import globals from "globals";
import pluginJs from "@eslint/js";
import tseslint from "typescript-eslint";
import pluginReact from "eslint-plugin-react";
import pluginPrettier from "eslint-plugin-prettier";
import pluginA11y from "eslint-plugin-jsx-a11y";
import pluginImport from "eslint-plugin-import";

/** @type {import('eslint').Linter.FlatConfig[]} */
export default [
    {
        files: ["**/*.{js,mjs,cjs,ts,jsx,tsx}"],
        languageOptions: {
            globals: globals.browser,
            ecmaVersion: "latest",
            sourceType: "module",
            parser: tseslint.parser
        },
        plugins: {
            react: pluginReact,
            "@typescript-eslint": tseslint.plugin,
            prettier: pluginPrettier,
            "jsx-a11y": pluginA11y,
            import: pluginImport,
        },
        rules: {
            ...pluginJs.configs.recommended.rules,
            ...tseslint.configs.recommended.rules,
            ...pluginReact.configs.recommended.rules,
            "prettier/prettier": "error",
            "react/display-name": "off",
            "jsx-a11y/anchor-is-valid": "off",
            "jsx-a11y/label-has-for": "off",
            "camelcase": "off",
            "func-names": ["error", "never"],
            "import/prefer-default-export": "off",
            "import/no-anonymous-default-export": "off",
            "import/no-extraneous-dependencies": "off",
            "no-multi-spaces": "off",
            "class-methods-use-this": "off",
            "no-class-assign": "off",
            "key-spacing": "off",
            "lines-between-class-members": "off",
            "no-param-reassign": "off",
            "consistent-return": "off",
            "jsx-a11y/href-no-hash": "off",
            "import/no-unresolved": "off",
            "no-tabs": "off",
            "react/react-in-jsx-scope": "off",
            "no-use-before-define": "off",
            "@typescript-eslint/no-use-before-define": "error",
            "react/jsx-filename-extension": ["error", { "extensions": [".tsx"] }],
            "react/prop-types": "off"
        },
    },
];