describe("Home Spec", () => {

    beforeEach(() => {
        cy.visit("/")
    })
    it("should have a title", () => {
        cy.get("title").should("contain", "World | The Daily Breif")
    })

    it("should have an h1", () => {
        cy.get("h1").should("contain", "World News")
    })

    it("should have the news field", () => {
        cy.get('[data-cy="news"]', { timeout: 10000 }).should("be.visible")

        cy.get('[data-cy="news-block"]').should("be.visible")
        cy.get('[data-cy="news-link"]').should("be.visible")
        cy.get('[data-cy="news-title"]').should("be.visible")
        cy.get('[data-cy="news-desc"]').should("be.visible")
    })
})
