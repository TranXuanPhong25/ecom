package com.ecom.productcategory.entities;

import jakarta.persistence.*;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Entity
@Data
@Table(name = "product_category_closure")
@NoArgsConstructor
@AllArgsConstructor
@IdClass(ProductCategoryClosureId.class)
public class ProductCategoryClosureEntity {
    @Id
    @Column(name = "ancestor_id")
    private Integer ancestorId;

    @Id
    @Column(name = "descendant_id")
    private Integer descendantId;

    @Column(name = "depth")
    private Integer depth;



}
